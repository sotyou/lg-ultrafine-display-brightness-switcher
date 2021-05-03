<template>
  <v-app id="inspire" color="white">
    <v-container fluid class="px-0">
      <v-layout>
        <v-flex xs12 sm6>
          <v-img
            :src="require('./assets/images/lg-ultrafine-24.png')"
            alt="Sunny image"
            width="400"
          ></v-img>
        </v-flex>
      </v-layout>
    </v-container>
    <v-card
      class="mx-auto"
      elevation="20"
      max-width="400"
      max-height="200"
      color="#E0F7FA"
    >
      <v-toolbar
        flat
        dense
        color="#78909C"
      >
        <v-toolbar-title>
          <span class="subheading"><h3>Brightness Slider<h1></span>
        </v-toolbar-title>
        <v-spacer></v-spacer>
      </v-toolbar>
      <v-card-text>
        <v-row
          class="mb-4"
          justify="space-between"
        >
          <v-col class="text-left">
            <span
              class="display-3 font-weight-light"
              v-text="brightness"
            ></span>
            <span class="subheading font-weight-light mr-1">Percent</span>
          </v-col>
          <v-col class="text-right">
            <v-btn
              :color="color"
              dark
              depressed
              fab
            >
              <v-icon large>
                mdi-brightness-6
              </v-icon>
            </v-btn>
          </v-col>
        </v-row>
        <v-slider
          v-model="brightness"
          :color="color"
          track-color="grey"
          always-dirty
          min="1"
          max="100"
        >
          <template v-slot:prepend>
            <v-icon
              :color="color"
              @click="decrement"
            >
              mdi-minus
            </v-icon>
          </template>
          <template v-slot:append>
            <v-icon
              :color="color"
              @click="increment"
            >
              mdi-plus
            </v-icon>
          </template>
        </v-slider>
      </v-card-text>
    </v-card>
  </v-app>
</template>

<script>

export default {
  mounted: function() {
    window.backend.basic(0).then(result => {
      this.init = true
      this.brightness = result
    })
  },

  data () {
    return {
      brightness: 0,
      init: false
    }
  },

  methods: {
    setBrightness: function() {
      window.backend.basic(this.brightness)
    },
    decrement () {
      this.brightness--
    },
    increment () {
      this.brightness++
    },
  },

  computed: {
    color () {
      if (this.brightness < 25) {
        return 'indigo'
      }
      if (this.brightness < 50) {
        return 'teal'
      }
      if (this.brightness < 75) {
        return 'green'
      }
      if (this.brightness < 100) {
        return 'orange'
      }
      return 'red'
    },
  },

  watch: {
    brightness: function(val) {
      if (this.init) {
        window.backend.basic(val)
      }
    }
  }
};
</script>

<style>
  @keyframes metronome-example {
    from {
      transform: scale(.5);
    }

    to {
      transform: scale(1);
    }
  }

  h3 {
    color: #FF9800
  }

  .v-avatar--metronome {
    animation-name: metronome-example;
    animation-iteration-count: infinite;
    animation-direction: alternate;
  }

  /* .logo {
    width: 16em;
  } */
</style>
