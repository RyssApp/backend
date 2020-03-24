package app.ryss.gateway.ratelimits

import io.ktor.application.ApplicationCall
import io.ktor.util.pipeline.PipelineContext
import java.time.Duration
import java.time.Instant
import java.time.temporal.ChronoUnit

/**
 * Implementation of [RateLimiter] that stores rate-limits in memory.
 */
class MemoryRateLimiter : RateLimiter {

    private val storage = mutableMapOf<String, MutableMap<String, Call>>()

    override fun isRateLimited(
        context: PipelineContext<*, ApplicationCall>,
        identifier: String,
        maxRequests: Int,
        maxTime: Long,
        maxUnit: ChronoUnit,
        action: String
    ): RateLimiter.Result {
        val requests = storage.computeIfAbsent(action) {
            mutableMapOf()
        }.computeIfAbsent(identifier) {
            Call(0, Instant.now())
        }
        requests.count++

        val remaining = (maxRequests - requests.count).coerceAtLeast(0)
        val reset = requests.createdAt.plus(maxTime, maxUnit).toEpochMilli()

        if (requests.count > maxRequests) {
            if (Duration.between(requests.createdAt, Instant.now()) < Duration.of(maxTime, maxUnit)) {
                return RateLimiter.Result(maxRequests, remaining, reset, true)
            } else {
                storage.remove(identifier)
            }
        } else {
            storage[identifier]?.replace(identifier, requests)
        }

        return RateLimiter.Result(maxRequests, remaining, reset, false)
    }

    private data class Call(var count: Int, val createdAt: Instant)

}