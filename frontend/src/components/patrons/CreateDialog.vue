<template>
  <div>
    <v-btn @click="show = true" class="ma-2 bg-primary">New patron</v-btn>
    <v-dialog v-model="show" width="33%">
      <v-card title="Add a new patron" class="ma-3 pa-5">
        <v-form v-model="valid">
          <v-text-field
            @update:focused="errResponse = {}"
            v-model="firstName"
            label="First name"
            :rules="[required]"
            :error-messages="firstNameErrors"
          />
          <v-text-field
            @update:focused="errResponse = {}"
            v-model="lastName"
            label="Last name"
            :rules="[required]"
            :error-messages="lastNameErrors"
          />

          <v-text-field
            @update:focused="errResponse = {}"
            v-model="weight"
            label="Weight"
            type="number"
            :rules="[required]"
            :error-messages="weightErrors"
          />
          <v-select
            @update:focused="errResponse = {}"
            v-model="sex"
            label="sex"
            :items="['M','F']"
            :rules="[required]"
            :error-messages="sexErrors"
          />
          <v-card-actions>
            <v-btn
              @click="submit"
              :disabled="!valid"
            >Submit</v-btn>
          </v-card-actions>
        </v-form>
      </v-card>
    </v-dialog>
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import { PatronsAPI } from '@/modules/api/patrons'
import { ApiError } from '@/modules/errors/ApiError'
import { required } from '@/modules/validation'

const api = new PatronsAPI();

const errResponse = ref<Record<string, string[]>>();
const valid = ref<boolean>(false);
const show = ref<boolean>(false);

const firstName = ref<string>("");
const lastName = ref<string>("");
const sex = ref<"M" | "F">("M");
const weight = ref<number>(80);

const firstNameErrors = computed(() => errResponse.value ?
  errResponse.value["firstName"] : []);
const lastNameErrors = computed(() => errResponse.value ?
  errResponse.value["lastName"] : []);
const sexErrors = computed(() => errResponse.value ?
  errResponse.value["sex"] : []);
const weightErrors = computed(() => errResponse.value ?
  errResponse.value["weight"] : []);

function submit() {
  if (valid.value) {
    api.CreatePatron({
      firstName: firstName.value,
      lastName: lastName.value,
      weight: weight.value.valueOf(),
      sex: sex.value
    }).catch((e) => {
      if (e instanceof ApiError) {
        errResponse.value = e.response["Detail"] as unknown as Record<string, string[]>;
      }
    }).then(() => show.value = false);
  }
}
</script>

<style scoped>

</style>