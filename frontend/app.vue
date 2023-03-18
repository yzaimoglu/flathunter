<template>
  <div>
    <Html class="bg-white dark:bg-gray-900"></Html>
    <NuxtLayout>
      <NuxtPage />
    </NuxtLayout>
  </div>
</template>

<script setup>
  useHead({
    link: [
      {
        rel: "stylesheet",
        href: "/styles/tailwind.css",
      },
    ],
  });
</script>