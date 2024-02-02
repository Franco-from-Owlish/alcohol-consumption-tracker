<template>
  <CreateDialog/>
  <PatronItem
    v-for="patron in patrons"
    :key="patron.id"
    :patron="patron"
  />
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { PatronsAPI } from '@/modules/api/patrons'
import PatronItem from '@/components/patrons/PatronItem.vue'
import { usePatronStore } from '@/stores/patrons'
import CreateDialog from '@/components/patrons/CreateDialog.vue'

const patronStore = usePatronStore();

const patrons = computed(() => patronStore.patrons);

const api = new PatronsAPI()

function load() {
  api.GetAll();
}

load();
</script>
