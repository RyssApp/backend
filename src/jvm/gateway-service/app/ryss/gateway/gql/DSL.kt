package app.ryss.gateway.gql

import graphql.Scalars
import graphql.schema.*

/**
 * Adds a field to the graphql object.
 */
fun GraphQLObjectType.Builder.addField(
    name: String,
    type: GraphQLType,
    block: GraphQLFieldDefinition.Builder.() -> GraphQLFieldDefinition.Builder = { this }
): GraphQLObjectType.Builder =
    field(GraphQLFieldDefinition.newFieldDefinition().apply {
        require(type is GraphQLOutputType) { "Type has to inherit GraphQLOutputType" }
        name(name)
        type(type)
        block(this)
    }.build())

/**
 * Adds an argument to the graphql field.
 */
fun GraphQLFieldDefinition.Builder.addArgument(
    name: String,
    type: GraphQLType,
    block: GraphQLArgument.Builder.() -> GraphQLArgument.Builder = { this }
): GraphQLFieldDefinition.Builder {
    return argument(GraphQLArgument.newArgument().apply {
        require(type is GraphQLInputType) { "Type has to inherit GraphQLInputType" }
        name(name)
        type(type)
        block(this)
    }.build())
}

/**
 * Returns the corresponding GQL type.
 */
fun String.Companion.toGraphQLType(notNull: Boolean = true): GraphQLType =
    graphQLType(Scalars.GraphQLString, notNull)

/**
 * Returns the corresponding GQL type.
 */
fun Int.Companion.toGraphQLType(notNull: Boolean = true): GraphQLType = graphQLType(Scalars.GraphQLInt, notNull)

/**
 * Returns the corresponding GQL type.
 */
fun Boolean.Companion.toGraphQLType(notNull: Boolean = true): GraphQLType =
    graphQLType(Scalars.GraphQLBoolean, notNull)

/**
 * Returns the corresponding GQL type.
 */
fun Float.Companion.toGraphQLType(notNull: Boolean = true): GraphQLType =
    graphQLType(Scalars.GraphQLFloat, notNull)

private fun graphQLType(type: GraphQLType, notNull: Boolean): GraphQLType =
    if (notNull) GraphQLNonNull.nonNull(type) else type
