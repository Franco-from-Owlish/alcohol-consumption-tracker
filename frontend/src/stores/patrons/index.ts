import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import type { Patron } from '@/modules/api/patrons/types'

export const usePatronStore = defineStore('patron', () => {
  const patrons = ref<Patron[]>()

  function removePatron(id: string) {
    const idx = patrons.value?.findIndex(x => x.id === id)
    if (idx && (idx >= 0)) {
      patrons.value?.splice(idx, 1)
    }
  }

  return { patrons, removePatron }
})
