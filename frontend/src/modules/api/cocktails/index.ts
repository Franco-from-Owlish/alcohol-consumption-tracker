import type { Cocktail } from '@/modules/api/cocktails/types'
import type { AxiosInstance, AxiosResponse } from 'axios'
import {api as defaultAPI} from '@/modules/axios'
import { AxiosError } from 'axios'
import { ApiError } from '@/modules/errors/ApiError'

export class CocktailsAPI {
  private api: AxiosInstance;

  constructor(api?: AxiosInstance) {
    this.api = api ?? defaultAPI;
  }

  async GetRandom(): Promise<Cocktail | null> {
    let response: AxiosResponse<Cocktail, any>;
    try {
      response = await this.api.get<Cocktail>("/cocktail/random")
    } catch (e) {
      if (e instanceof AxiosError) {
        throw new ApiError({
          name: 'GET_RANDOM_COCKTAIL',
          message: e.message,
          cause: e
        })
      }
      throw e;
    }
    return response.data;
  }
}