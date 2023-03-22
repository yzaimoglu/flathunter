<template>
    <div>
        <div>
            <Header />
            <div class="flex flex-col justify-center min-h-screen bg-white dark:bg-gray-900">
                <slot />
                <Footer class="sticky top-[100vh]" />
            </div>
        </div>
        <CookieNotice />
    </div>
</template>

<script setup>

</script>
