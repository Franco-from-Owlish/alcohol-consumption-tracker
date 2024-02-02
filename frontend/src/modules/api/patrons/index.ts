import type { AxiosInstance, AxiosResponse } from 'axios'
import {api as defaultAPI} from '@/modules/axios'
import { AxiosError } from 'axios'
import { ApiError } from '@/modules/errors/ApiError'
import type { Patron } from '@/modules/api/patrons/types'
import { usePatronStore } from '@/stores/patrons'

export class PatronsAPI {
  private api: AxiosInstance;
  private store = usePatronStore();

  constructor(api?: AxiosInstance) {
    this.api = api ?? defaultAPI;
  }

  async GetAll(): Promise<Patron[] | null> {
    let response: AxiosResponse<Patron[], any>;
    try {
      response = await this.api.get<Patron[]>("/patron/")
    } catch (e) {
      if (e instanceof AxiosError) {
        throw new ApiError({
          name: 'GET_ALL_PATRONS',
          message: e.message,
          cause: e
        })
      }
      throw e;
    }
    this.store.$patch({
      patrons: response.data
    })
    return response.data;
  }

  async CreatePatron(patron: Partial<Patron>): Promise<Patron| null> {
    let response: AxiosResponse<Patron, any>;
    try {
      response = await this.api.post<Patron>("/patron/", {
        ...patron,
        weight: parseFloat(<string>(patron.weight ?? "0"))
      })
    } catch (e) {
      if (e instanceof AxiosError) {
        throw new ApiError({
          name: 'GET_ALL_PATRONS',
          message: e.message,
          cause: e
        })
      }
      throw e;
    }
    this.GetAll().then();
    return response.data;
  }

  async GetPatron(id: string): Promise<Patron[] | null> {
    let response: AxiosResponse<Patron[], any>;
    try {
      response = await this.api.get<Patron[]>(`/patron/${id}`);
    } catch (e) {
      if (e instanceof AxiosError) {
        throw new ApiError({
          name: 'GET_PATRON',
          message: e.message,
          cause: e
        })
      }
      throw e;
    }
    return response.data;
  }

  async DeletePatron(id: string): Promise<void> {
    try {
      await this.api.delete<Patron[]>(`/patron/${id}`)
    } catch (e) {
      if (e instanceof AxiosError) {
        throw new ApiError({
          name: 'DELETE_PATRON',
          message: e.message,
          cause: e
        })
      }
      throw e;
    }
    this.store.removePatron(id);
  }

}