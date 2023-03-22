<template>
    <div class="relative flex justify-center">
        <div v-if="modalState" class="absolute inset-0 z-10 w-full h-full bg-gray-900 opacity-90"></div>
        <transition
            :enter-from-class="'opacity-0'"
            :leave-to-class="'opacity-0'"
            :enter-active-class="'transition duration-500 ease-in'"
            :leave-active-class="'transition duration-500 ease-in'"
        >
            <div v-if="modalState" class="fixed inset-0 z-20 overflow-y-auto">
                <div class="flex items-end justify-center min-h-screen px-4 pt-4 pb-20 text-center sm:block sm:p-0">
                    <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>

                    <div class="relative inline-block p-4 overflow-hidden text-left align-middle transition-all transform bg-white shadow-xl sm:max-w-lg rounded-xl dark:bg-gray-900 sm:my-8 sm:w-full sm:p-6">
                        <div class="flex items-center justify-center mx-auto">
                            <iframe class="h-full rounded-lg" width="3424" height="1167" src="https://www.youtube.com/embed/dQw4w9WgXcQ" title="Rick Astley - Never Gonna Give You Up (Official Music Video)" frameborder="0"></iframe>
                        </div>

                        <div class="mt-5 text-center">
                            <h3 class="text-lg font-medium text-gray-800 dark:text-white" id="modal-title">
                                Get Rickrolled Bitch
                            </h3>
                        </div>

                        <div class="mt-4 sm:flex sm:items-center sm:justify-between sm:mt-6 sm:-mx-2">
                            <button @click="() => modalState = false" class="px-4 sm:mx-2 w-full py-2.5 text-sm font-medium dark:text-gray-200 dark:border-gray-700 dark:hover:bg-gray-800 tracking-wide text-gray-700 capitalize transition-colors duration-300 transform border border-gray-200 rounded-md hover:bg-gray-100 focus:outline-none focus:ring focus:ring-gray-300 focus:ring-opacity-40">
                                Schließen
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </transition>
    </div>
</template>

<script setup>
    const modalState = useState('indexViewDemoModalState', () => false)
</script>
