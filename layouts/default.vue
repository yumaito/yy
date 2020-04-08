<template>
  <v-app>
    <v-navigation-drawer
      v-model="drawer"
      clipped
      fixed
      app
    >
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="title">
            メニュー
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-divider />
      <v-list
        dense
        nav
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
      clipped-left
      flat
      dark
      fixed
      app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-btn
        icon
        nuxt
        to="/"
      >
        <v-icon>mdi-application</v-icon>
      </v-btn>
      <v-spacer />
      <v-btn
        text
        :to="{'name': 'login'}"
      >
        <v-icon>mdi-logout-variant</v-icon>
        Logout
      </v-btn>
    </v-app-bar>
    <v-content>
      <v-container>
        <nuxt />
      </v-container>
    </v-content>
    <v-footer
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
      data
    }
  }
}
</script>
