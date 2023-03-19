<template>
    <div v-cloak>
        <LoadingSpinner v-if="!pageLoaded" />
        <transition 
            :enter-from-class="'opacity-0'"
            :leave-to-class="'opacity-0'"
            :enter-active-class="'transition duration-500 ease-in'"
            :leave-active-class="'transition duration-500 ease-out'"
        >
            <div v-if="pageLoaded">
                <DashboardOptions class="container mx-auto py-10 px-0" />
                <DashboardDivider />
                <DashboardFlatCards class="container mx-auto py-20 px-0" />
            </div>
        </transition>
    </div>
</template>

<script setup>
    const pageLoaded = ref(false)

    onMounted(() => {
        setTimeout(() => {
            pageLoaded.value = true
        }, 200)
    })
</script>