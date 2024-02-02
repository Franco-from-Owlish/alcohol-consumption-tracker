import type { AxiosError } from 'axios'

type ErrorName =
  | 'GET_RANDOM_COCKTAIL'
  | 'CREATE_PATRON'
  | 'GET_ALL_PATRONS'
  | 'GET_PATRON'
  | 'DELETE_PATRON';

export class ApiError extends Error {
  name: ErrorName
  message: string
  response: Record<string,string|string[]>
  cause: AxiosError

  constructor(
    { name, message, cause }:
      { name: ErrorName, message: string, cause: AxiosError }
  ) {
    super();
    this.name = name;
    this.message = message;
    this.response = cause.response?.data as Record<string,string|string[]>;
    this.cause = cause;
  }
}

