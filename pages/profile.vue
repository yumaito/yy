<template>
  <v-layout>
    <h2>{{ user.name }}</h2>
    <p>{{ user }}</p>
  </v-layout>
</template>

<script>
export default {
  name: 'Profile',
  computed: {
    user () {
      return this.$auth.user
    }
  }
}
</script>
