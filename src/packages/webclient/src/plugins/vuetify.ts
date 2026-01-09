/**
 * plugins/vuetify.ts
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'
import {md3} from 'vuetify/blueprints'

// Composables
import {createVuetify, type ThemeDefinition} from 'vuetify'

const darkTheme: ThemeDefinition = {
  dark: true,
  colors: {
    // Primary (purple accent)
    primary: '#C9A7FF',
    'on-primary': '#2A004F',
    'primary-container': '#3E1A6D',
    'on-primary-container': '#E9DDFF',

    // Secondary (muted purple / gray)
    secondary: '#B9A6C9',
    'on-secondary': '#2C1E35',
    'secondary-container': '#43314F',
    'on-secondary-container': '#EADCF4',

    // Tertiary (slightly warmer violet)
    tertiary: '#E0B7FF',
    'on-tertiary': '#3A004D',
    'tertiary-container': '#552068',
    'on-tertiary-container': '#F3DBFF',

    // Error (unchanged but tuned darker)
    error: '#FFB4AB',
    'on-error': '#690005',
    'error-container': '#93000A',
    'on-error-container': '#FFDAD6',

    // Backgrounds (near-black with purple bias)
    background: '#0F0B14',
    'on-background': '#E6DFF0',

    surface: '#0F0B14',
    'on-surface': '#E6DFF0',
    'surface-variant': '#2A2233',
    'on-surface-variant': '#CCC2D9',

    // Utility
    outline: '#8C8299',
    shadow: '#000000',

    // Inverse (for light surfaces like dialogs/snackbars)
    'inverse-surface': '#E6DFF0',
    'inverse-on-surface': '#1E1824',
    'inverse-primary': '#5E2D91',
  },
}

export default createVuetify({
  blueprint: md3,
  theme: {
    defaultTheme: 'darkTheme',
    themes: {
      darkTheme
    }
  },
})
