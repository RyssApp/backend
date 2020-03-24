package app.ryss.gateway.gql

import graphql.schema.GraphQLFieldDefinition
import graphql.schema.GraphQLObjectType

/**
 * GraphQL user type.
 */
val userType: GraphQLObjectType = GraphQLObjectType.newObject()
    .name("User")
    .addField("id", String.toGraphQLType())
    .addField("email", String.toGraphQLType())
    .addField("verified", Boolean.toGraphQLType())
    .addField("flags", Int.toGraphQLType())
    .addField("avatar", String.toGraphQLType(notNull = false))
    .addField("locale", String.toGraphQLType())
    .addField("name", String.toGraphQLType(notNull = false))
    .addField("disable", Boolean.toGraphQLType())
    .addField("created_at", String.toGraphQLType())
    .addField("last_login", String.toGraphQLType(notNull = false))
    .build()

/**
 * GraphQL login response type.
 */
val loginResponseType: GraphQLObjectType = GraphQLObjectType.newObject()
    .name("LoginResponse")
    .addField("user", userType)
    .addField("token", String.toGraphQLType())
    .build()

/**
 * GraphQL register mutation.
 */
val registerMutation: GraphQLFieldDefinition = GraphQLFieldDefinition.newFieldDefinition()
    .type(userType)
    .name("register")
    .addArgument("email", String.toGraphQLType())
    .addArgument("password", String.toGraphQLType())
    .build()

/**
 * GraphQL login mutation.
 */
val loginMutation: GraphQLFieldDefinition = GraphQLFieldDefinition.newFieldDefinition()
    .type(loginResponseType)
    .name("login")
    .addArgument("login", String.toGraphQLType())
    .addArgument("password", String.toGraphQLType())
    .build()
