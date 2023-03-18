<template>
    <transition 
        :enter-from-class="'opacity-0'"
        :leave-to-class="'opacity-0'"
        :enter-active-class="'transition duration-700 fade-in'"
        :leave-active-class="'transition duration-700 fade-out'"
    >
        <section v-if="!cookieStatus" class="fixed max-w-md p-4 mx-auto bg-white border border-gray-200 dark:bg-gray-800 left-12 bottom-16 dark:border-gray-700 rounded-2xl">
            <h2 class="font-semibold text-gray-800 dark:text-white">🍪 Cookie Notice</h2>

            <p class="mt-4 text-sm text-gray-600 dark:text-gray-300">We use cookies to ensure that we give you the best experience on our website. <a href="#" class="text-blue-500 hover:underline">Read cookies policies</a>. </p>
            
            <div class="flex items-center justify-between mt-4 gap-x-4 shrink-0">
                <button class="text-xs text-gray-800 underline transition-colors duration-300 dark:text-white dark:hover:text-gray-400 hover:text-gray-600 focus:outline-none">
                    Manage your preferences
                </button>

                <button @click=acceptCookie class=" text-xs bg-gray-900 font-medium rounded-lg hover:bg-gray-700 text-white px-4 py-2.5 duration-300 transition-colors focus:outline-none">
                    Accept
                </button>
            </div>
        </section>
    </transition>
</template>

<script setup>
    const cookieStatus = useState('cookieStatus', () => false)

    const acceptCookie = () => {
        cookieStatus.value = true
    }
</script>