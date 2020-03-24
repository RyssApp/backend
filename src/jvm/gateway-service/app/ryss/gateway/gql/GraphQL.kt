package app.ryss.gateway.gql

import app.ryss.gateway.core.GatewayService
import graphql.schema.FieldCoordinates
import graphql.schema.GraphQLCodeRegistry
import graphql.schema.GraphQLObjectType
import graphql.schema.GraphQLSchema
import io.ktor.application.ApplicationCall
import io.ktor.routing.Routing
import io.ktor.util.pipeline.PipelineContext
import ktor.graphql.config
import ktor.graphql.graphQL

/**
 * GraphQL manager.
 * @param enableGraphiQl whether graphiql should be enabled or not
 * @param application gateway-service instance
 */
class GraphQL(private val enableGraphiQl: Boolean, private val application: GatewayService) {

    /**
     * Registers the /graphql route to [routing].
     */
    operator fun invoke(routing: Routing): Routing = routing.apply(::registerRoute)

    private fun registerRoute(routing: Routing) {
        routing.apply {
            graphQL("/graphql", buildSchema()) {
                config {
                    graphiql = enableGraphiQl
                    context = RequestContext(this@graphQL, this@GraphQL.application)
                }
            }
        }
    }
}

/**
 * Context of a request.
 * @property application gateway-service instance
 */
class RequestContext(pipelineContext: PipelineContext<Unit, ApplicationCall>, val application: GatewayService) :
    PipelineContext<Unit, ApplicationCall> by pipelineContext {
    /**
     * @see PipelineContext
     */
    operator fun component1(): RequestContext = this

    /**
     * @see application
     */
    operator fun component2(): GatewayService = application
}


private fun buildSchema(): GraphQLSchema {
    return GraphQLSchema.newSchema()
        .query(createQuery())
        .mutation(createMutation())
        .codeRegistry(buildRuntimeWiring())
        .build()
}

private fun buildRuntimeWiring(): GraphQLCodeRegistry {
    return GraphQLCodeRegistry.newCodeRegistry()
        .dataFetcher(
            FieldCoordinates.coordinates("Mutation", "register"),
            ratelimitedDataFetcher(5, name = "register") {
                return@ratelimitedDataFetcher object {
                    val name: String = "Test"
                    val id: String = "Test"
                }
            })
        .build()
}

private fun createQuery(): GraphQLObjectType {
    return GraphQLObjectType.newObject()
        .name("Query")
        .field(storeQuery)
        .build()
}

private fun createMutation(): GraphQLObjectType {
    return GraphQLObjectType.newObject()
        .name("Mutation")
        .field(registerMutation)
        .field(loginMutation)
        .build()
}
