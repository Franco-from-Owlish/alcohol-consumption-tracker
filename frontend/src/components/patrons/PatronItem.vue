<template>
  <v-card class="bg-transparent elevation-2 ma-2">
    <v-row>
      <v-col cols="4">
        <v-card-title>{{ patron.firstName }} {{ patron.lastName }}</v-card-title>
        <v-card-subtitle>id: {{ patron.id }}</v-card-subtitle>
      </v-col>
      <v-col cols="4">
        <v-card-text>
          <v-list class="bg-transparent">
            <v-list-item>
              Sex: {{ patron.sex }}
            </v-list-item>
            <v-list-item>
              Weight: {{ patron.weight }}
            </v-list-item>
            <v-list-item>
              EBAC: {{ Math.round((patron.ebac || 0) * 1000)/1000 }} %
            </v-list-item>
          </v-list>
        </v-card-text>
      </v-col>
      <v-col cols="4">
        <v-card-actions>
          <v-btn @click="api.DeletePatron(patron.id as string)">
            Remove patron
          </v-btn>
        </v-card-actions>
      </v-col>
    </v-row>
  </v-card>
</template>

<script setup lang="ts">
import type { Patron } from '@/modules/api/patrons/types'
import { PatronsAPI } from '@/modules/api/patrons'

defineProps<{
  patron: Patron;
}>()

const api = new PatronsAPI();
</script>