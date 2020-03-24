package app.ryss.gateway.gql

import app.ryss.gateway.core.GatewayService
import app.ryss.gateway.ratelimits.RateLimiter
import graphql.GraphqlErrorBuilder
import graphql.execution.DataFetcherResult
import graphql.schema.DataFetcher
import graphql.schema.DataFetchingEnvironment
import io.ktor.application.ApplicationCall
import io.ktor.application.call
import io.ktor.util.pipeline.PipelineContext
import kotlinx.coroutines.runBlocking
import java.time.temporal.ChronoUnit


/**
 * Creates an [RatelimitedDataFetcher] with [maxRequests] and [name].
 * @param maxRequests maximal amount of requests
 * @param maxTime amount of time until reset
 * @param maxUnit [ChronoUnit] of [maxTime]
 * @param name name of the request path
 * @param block the data fetcher
 */
fun <T : Any> ratelimitedDataFetcher(
    maxRequests: Int = RateLimiter.MAX_REQUESTS,
    maxTime: Long = 1,
    maxUnit: ChronoUnit = ChronoUnit.MINUTES,
    name: String,
    block: RatelimitedDataFetcher.DataFetchContext.() -> T
): RatelimitedDataFetcher<T> {
    return object : RatelimitedDataFetcher<T> {
        override val maxRequests: Int
            get() = maxRequests

        override val name: String
            get() = name

        override val maxTime: Long
            get() = maxTime

        override val maxUnit: ChronoUnit
            get() = maxUnit

        override fun fetch(context: RatelimitedDataFetcher.DataFetchContext): T = block(context)
    }
}

/**
 * Implementation of [DataFetcher] that uses [GatewayService.rateLimiter] to rate-limit requests.
 */
interface RatelimitedDataFetcher<T : Any> : DataFetcher<Any> {

    /**
     * Maximal amount of requests
     */
    val maxRequests: Int

    /**
     * @see RateLimiter
     */
    val maxTime: Long

    /**
     * @see RateLimiter
     */
    val maxUnit: ChronoUnit

    /**
     * Name of the request
     */
    val name: String

    /**
     * Rate-limits requests before
     * @see DataFetcher.get
     */
    override fun get(environment: DataFetchingEnvironment): Any {
        val context = environment.getContext<RequestContext>()
        val (pipelineContext, application) = context

        val limit = application.rateLimiter.isRateLimited(pipelineContext, maxRequests, maxTime, maxUnit, name)

        runBlocking {
            RateLimiter.addHeaders(limit, pipelineContext.call.response)
        }

        if (limit.ratelimited) {
            return DataFetcherResult.newResult<T>()
                .error(
                    GraphqlErrorBuilder.newError()
                        .message("You're being ratelimited")
                        .build()
                )
                .build()
        }

        return fetch(DataFetchContext(context, environment))
    }

    /**
     * Fetches the data.
     */
    fun fetch(
        context: DataFetchContext
    ): T

    /**
     * Context for graphql data fetchers.
     * @property requestContext the [RequestContext] of the request
     * @property environment the [DataFetchingEnvironment]
     */
    data class DataFetchContext(val requestContext: RequestContext, val environment: DataFetchingEnvironment) {

        /**
         * Convenience getter.
         */
        val pipelineContext: PipelineContext<*, ApplicationCall>
            get() = requestContext

        /**
         * [GatewayService] instance.
         */
        val application: GatewayService
            get() = requestContext.application
    }
}