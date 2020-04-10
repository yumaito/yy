<template>
  <v-app>
    <v-navigation-drawer
      v-model="drawer"
      :clipped="clipped"
      fixed
      app
    >
      <v-list
        dense
      >
        <template v-for="item in data.menues">
          <!-- サブメニューありの場合 -->
          <v-list-group
            v-if="item.lists"
            :key="item.title"
            :prepend-icon="item.icon"
            no-action
            :append-icon="item.lists ? undefined : ''"
          >
            <template v-slot:activator>
              <v-list-item-content>
                <v-list-item-title>{{ item.title }}</v-list-item-title>
              </v-list-item-content>
            </template>
            <v-list-item
              v-for="list in item.lists"
              :key="list.title"
              nuxt
              :to="list.to"
            >
              <v-list-item-title>{{ list.title }}</v-list-item-title>
              <v-list-item-icon>
                <v-icon>{{ list.icon }}</v-icon>
              </v-list-item-icon>
            </v-list-item>
          </v-list-group>
          <!-- サブメニューなしの場合 -->
          <v-list-item
            v-else
            :key="item.title"
            nuxt
            :to="item.to"
          >
            <v-list-item-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-item-action>
            <v-list-item-content>
              <v-list-item-title>
                {{ item.title }}
              </v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </template>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar
      :clipped-left="clipped"
      :color="color"
      dark
      fixed
      app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-btn
        to="/"
        nuxt
        icon
      >
        <v-icon>mdi-application</v-icon>
      </v-btn>
      <v-spacer />
      <v-btn
        :to="{'name': 'user'}"
        nuxt
        icon
      >
        <v-icon>mdi-account</v-icon>
      </v-btn>
    </v-app-bar>
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>
    <v-footer
      :color="color"
      dark
      fixed
      app
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>

<script>
import data from '@/data'

export default {
  data () {
    return {
      drawer: null,
      color: 'primary',
      clipped: this.$vuetify.breakpoint.lgAndUp,
      data
    }
  }
}
</script>
