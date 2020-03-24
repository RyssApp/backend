package app.ryss.gateway.gql

import graphql.schema.GraphQLFieldDefinition
import graphql.schema.GraphQLObjectType

/**
 * Graphql store type.
 */
val storeType: GraphQLObjectType = GraphQLObjectType.newObject()
    .name("Store")
    .addField("id", String.toGraphQLType())
    .addField("display_name", String.toGraphQLType())
    .addField("logo", String.toGraphQLType())
    .addField("location", String.toGraphQLType())
    .addField("created_at", String.toGraphQLType())
    .build()

/**
 * Graphql store query.
 */
val storeQuery: GraphQLFieldDefinition = GraphQLFieldDefinition.newFieldDefinition()
    .type(storeType)
    .name("store")
    .addArgument("id", String.toGraphQLType())
    .build()
