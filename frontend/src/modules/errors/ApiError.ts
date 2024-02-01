import type { AxiosError } from 'axios'

type ErrorName =
  | 'GET_RANDOM_COCKTAIL';

export class ApiError extends Error {
  name: ErrorName
  message: string
  response: object
  cause: AxiosError

  constructor(
    { name, message, cause }:
      { name: ErrorName, message: string, cause: AxiosError }
  ) {
    super();
    this.name = name;
    this.message = message;
    this.response = cause.response?.data as object;
    this.cause = cause;
  }
}

