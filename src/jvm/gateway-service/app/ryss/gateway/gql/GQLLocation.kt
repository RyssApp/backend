package app.ryss.gateway.gql

import graphql.schema.GraphQLObjectType

/**
 * GraphQL location type.
 */
val locationType: GraphQLObjectType = GraphQLObjectType.newObject()
    .name("Location")
    .addField("address", String.toGraphQLType())
    .addField("city", String.toGraphQLType())
    .addField("zip_code", String.toGraphQLType())
    .addField("latitude", Float.toGraphQLType())
    .addField("longitude", Float.toGraphQLType())
    .build()