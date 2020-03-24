package app.ryss.gateway.ratelimits

import app.ryss.gateway.entites.RequestError
import app.ryss.gateway.util.hashSHA256
import io.ktor.application.ApplicationCall
import io.ktor.application.call
import io.ktor.features.origin
import io.ktor.http.HttpStatusCode
import io.ktor.response.ApplicationResponse
import io.ktor.response.header
import io.ktor.response.respond
import io.ktor.util.pipeline.PipelineContext
import java.time.temporal.ChronoUnit

/**
 * RateLimiter interface.
 */
interface RateLimiter {

    /**
     * Returns whether the request in [context] is rate limited or not.
     * @param identifier identifier of requester (normally IP-Hash)
     * @param maxRequests maximal amount of requessts
     * @param maxTime amount of time until reset
     * @param maxUnit [ChronoUnit] of [maxTime]
     * @param action action to be rate-limited
     */
    fun isRateLimited(
        context: PipelineContext<*, ApplicationCall>,
        identifier: String,
        maxRequests: Int = MAX_REQUESTS,
        maxTime: Long = 1,
        maxUnit: ChronoUnit = ChronoUnit.MINUTES,
        action: String = "/"
    ): Result

    /**
     * Calls [isRateLimited] with the SHA-256 hashed ip of [context].
     * @see isRateLimited
     * @see hashSHA256
     */
    fun isRateLimited(
        context: PipelineContext<*, ApplicationCall>,
        maxRequests: Int = MAX_REQUESTS,
        maxTime: Long = 1,
        maxUnit: ChronoUnit = ChronoUnit.MINUTES,
        action: String
    ): Result =
        isRateLimited(context, context.context.request.origin.host.hashSHA256(), maxRequests, maxTime, maxUnit, action)

    /**
     * Interceptor for [io.ktor.application.ApplicationCallPipeline].
     */
    suspend fun intercept(context: PipelineContext<*, ApplicationCall>) {
        val limit = isRateLimited(context, "/")

        if (limit.ratelimited) {
            addHeaders(limit, context.call.response)
            respond(context.call)
            context.finish()
        } else {
            context.proceed()
        }
    }

    /**
     * Rate-limiting result.
     * @property max maximal amount of requests
     * @property remaining remaining amount of requests
     * @property reset timestamp of reset
     * @property ratelimited whether the request was rate-limited or not
     */
    data class Result(val max: Int, val remaining: Int, val reset: Long, val ratelimited: Boolean)


    companion object {
        /**
         * The maximal amount of requests per minute.
         */
        const val MAX_REQUESTS: Int = 150

        /**
         * Adds rate-limit headers to [response] based on [limit].
         */
        fun addHeaders(limit: Result, response: ApplicationResponse) {
            response.header("X-RateLimit-Limit", limit.max)
            response.header("X-RateLimit-Remaining", limit.remaining)
            response.header("X-RateLimit-Reset", limit.reset)
        }

        /**
         * Responds to a [call] with an rate limit error message.
         */
        suspend fun respond(call: ApplicationCall) {
            call.respond(
                HttpStatusCode.TooManyRequests,
                RequestError(
                    "Ratelimit hit",
                    "You hit the ratelimit",
                    HttpStatusCode.TooManyRequests.value
                )
            )
        }
    }
}
