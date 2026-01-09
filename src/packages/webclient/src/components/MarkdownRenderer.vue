<script setup>
import {computed} from 'vue'
import MarkdownIt from 'markdown-it'
import DOMPurify from 'dompurify'

const md = new MarkdownIt({
  html: true,
  linkify: true,
  typographer: true
})

const props = defineProps({
  markdown: {
    type: String,
    default: ''
  }
})

const renderedHtml = computed(() => {
  const unsafeHtml = md.render(props.markdown)
  return DOMPurify.sanitize(unsafeHtml)
})
</script>

<template>
  <div class="markdown-content" v-html="renderedHtml"></div>
</template>

<style scoped>

::v-deep * {
  padding: revert;
  margin: revert;
}

::v-deep p {
  text-align: justify;
}

::v-deep img {
  max-width: 100%;
  height: auto;
  display: block;
  margin: 1em auto;
}
</style>
