<template>
  <v-row dense>
    <v-col
      v-for="item in showMenues"
      :key="item.title"
      :cols="4"
    >
      <v-card outlined>
        <v-card-title>
          <v-icon>{{ item.icon }}</v-icon>
          {{ item.title }}
        </v-card-title>
        <v-card-actions>
          <v-btn
            v-for="list of item.lists"
            :key="list.title"
            :to="list.to"
            nuxt
            text
          >
            {{ list.title }}
          </v-btn>
          <v-spacer />
          <v-btn
            v-if="!item.lists"
            :to="item.to"
            text
            nuxt
          >
            <v-icon>mdi-chevron-right</v-icon>
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script>
import data from '@/data'
export default {
  data () {
    return {
      data
    }
  },
  computed: {
    showMenues () {
      return this.data.menues.filter(ms => !ms.hide)
    }
  }
}
</script>
