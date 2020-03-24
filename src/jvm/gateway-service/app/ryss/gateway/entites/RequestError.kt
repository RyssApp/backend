package app.ryss.gateway.entites

/**
 * Representation of an error during requests.
 * @property name name of the error
 * @property message more meaningful explanation of the error
 * @property code http error code
 */
data class RequestError(val name: String?, val message: String?, val code: Int?)