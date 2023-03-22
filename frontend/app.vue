<template>
  <div>
    <div v-if="loaded">
      <Html :class='[!loaded ? "dark" : 0]' ></Html>
      <NuxtLayout>
        <NuxtPage class="dark:bg-gray-900" />
      </NuxtLayout>
    </div>
    <LoadingSpinner class="absolute right-1/2 bottom-1/2  transform translate-x-1/2 translate-y-1/2" v-else />
  </div>
</template>

<style>
  @import '/styles/custom.css';
</style>

<script setup>
  import { darkMode } from '/utilities/darkMode.js'

  const loaded = ref(false)
  const darkModePreference = useState('darkModePreference', () => true)

  onBeforeMount(() => {
    darkModePreference.value = localStorage.getItem('theme') === 'dark' ? true : false
    darkMode()
    setTimeout(() => {
      loaded.value = true
    }, 150)
  })

  window.onerror = function(e){
    if(e.includes("NotFoundError:")){
        document.location.reload()
        return true;
    }
  }

  useHead({
    link: [
      {
        rel: "stylesheet",
        href: "/styles/tailwind.css",
      },
    ],
  });
</script>
