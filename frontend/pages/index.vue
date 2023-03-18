<template>
    <div class="mt-5">
        <div v-if="modalState" class="absolute inset-0 z-10 w-full h-full bg-gray-900 opacity-90"></div>
        <transition
            :enter-from-class="'opacity-0'"
            :leave-to-class="'opacity-0'"
            :enter-active-class="'transition duration-300 fade-in'"
            :leave-active-class="'transition duration-300 fade-out'"
        >
            <section v-if="showUserLogin" class="max-w-4xl p-6 mx-auto bg-white rounded-md shadow-md dark:bg-gray-900">
                <h2 class="text-lg font-semibold text-gray-700 capitalize dark:text-white">Account settings</h2>
                <form>
                    <div class="grid grid-cols-1 gap-6 mt-4 sm:grid-cols-2">
                        <div>
                            <label class="text-gray-700 dark:text-gray-200" for="username">Username</label>
                            <input id="username" type="text" class="block w-full px-4 py-2 mt-2 text-gray-700 bg-white border border-gray-200 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-400 focus:ring-blue-300 focus:ring-opacity-40 dark:focus:border-blue-300 focus:outline-none focus:ring">
                        </div>

                        <div>
                            <label class="text-gray-700 dark:text-gray-200" for="emailAddress">Email Address</label>
                            <input id="emailAddress" type="email" class="block w-full px-4 py-2 mt-2 text-gray-700 bg-white border border-gray-200 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-400 focus:ring-blue-300 focus:ring-opacity-40 dark:focus:border-blue-300 focus:outline-none focus:ring">
                        </div>

                        <div>
                            <label class="text-gray-700 dark:text-gray-200" for="password">Password</label>
                            <input id="password" type="password" class="block w-full px-4 py-2 mt-2 text-gray-700 bg-white border border-gray-200 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-400 focus:ring-blue-300 focus:ring-opacity-40 dark:focus:border-blue-300 focus:outline-none focus:ring">
                        </div>

                        <div>
                            <label class="text-gray-700 dark:text-gray-200" for="passwordConfirmation">Password Confirmation</label>
                            <input id="passwordConfirmation" type="password" class="block w-full px-4 py-2 mt-2 text-gray-700 bg-white border border-gray-200 rounded-md dark:bg-gray-800 dark:text-gray-300 dark:border-gray-600 focus:border-blue-400 focus:ring-blue-300 focus:ring-opacity-40 dark:focus:border-blue-300 focus:outline-none focus:ring">
                        </div>
                    </div>

                    <div class="flex justify-end mt-6">
                        <button @click="accountSave" class="px-8 py-2.5 leading-5 text-white transition-colors duration-300 transform bg-gray-700 rounded-md hover:bg-gray-600 focus:outline-none focus:bg-gray-600">Save</button>
                    </div>
                </form>
            </section>
        </transition>
        <Faq />
        <div class="flex flex-col items-center w-full max-w-lg p-5 mx-auto mt-2 text-center bg-white dark:bg-gray-900">
            <button @click="() => showUserLogin = !showUserLogin" class="px-8 py-2.5 z-0 leading-5 text-white transition-colors duration-300 transform bg-gray-700 rounded-md hover:bg-gray-600 focus:outline-none focus:bg-gray-600">Toggle Account Settings</button>
        </div>
        <div>
            <label for="dropzone-file" class="flex flex-col items-center w-full max-w-lg p-5 mx-auto mt-2 text-center bg-white border-2 border-gray-300 border-dashed cursor-pointer dark:bg-gray-900 dark:border-gray-700 rounded-xl">
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-8 h-8 text-gray-500 dark:text-gray-400">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M12 16.5V9.75m0 0l3 3m-3-3l-3 3M6.75 19.5a4.5 4.5 0 01-1.41-8.775 5.25 5.25 0 0110.233-2.33 3 3 0 013.758 3.848A3.752 3.752 0 0118 19.5H6.75z" />
                </svg>
                <h2 class="mt-1 font-medium tracking-wide text-gray-700 dark:text-gray-200">Encrypted File</h2>
                <p class="mt-2 text-xs tracking-wide text-gray-500 dark:text-gray-400">Upload or drag & drop your file here.</p>
                <input id="dropzone-file" type="file" class="hidden" />
            </label>
        </div>
        <div class="relative flex justify-center">
            <button @click="() => { modalState = true }" class="px-6 z-0 py-2 mx-auto tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-md hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80">
                Open Modal
            </button>

            <!-- Modal -->
            <transition 
                :enter-from-class="'opacity-0'"
                :leave-to-class="'opacity-0'"
                :enter-active-class="'transition duration-500 ease-in'"
                :leave-active-class="'transition duration-500 ease-in'"
            >
                <div v-if="modalState" class="fixed inset-0 z-20 overflow-y-auto">
                    <div class="flex items-end justify-center min-h-screen px-4 pt-4 pb-20 text-center sm:block sm:p-0">
                        <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>

                        <div class="relative inline-block px-4 pt-5 pb-4 overflow-hidden text-left align-bottom transition-all transform bg-white rounded-lg shadow-xl rtl:text-right dark:bg-gray-900 sm:my-8 sm:align-middle sm:max-w-sm sm:w-full sm:p-6">
                            <div>
                                <div class="flex items-center justify-center">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="w-8 h-8 text-gray-700 dark:text-gray-300" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                        <path stroke-linecap="round" stroke-linejoin="round" d="M5 8h14M5 8a2 2 0 110-4h14a2 2 0 110 4M5 8v10a2 2 0 002 2h10a2 2 0 002-2V8m-9 4h4" />
                                    </svg>
                                </div>

                                <div class="mt-2 text-center">
                                    <h3 class="text-lg font-medium leading-6 text-gray-800 capitalize dark:text-white" id="modal-title">Archive Project</h3>
                                    <p class="mt-2 text-sm text-gray-500 dark:text-gray-400">
                                        Lorem, ipsum dolor sit amet consectetur
                                        adipisicing elit. Aspernatur dolorum aliquam ea, ratione deleniti porro officia? Explicabo
                                        maiores suscipit.
                                    </p>
                                </div>
                            </div>

                            <div class="mt-5 sm:flex sm:items-center sm:justify-between">
                                <a href="#" class="text-sm text-blue-500 hover:underline">Learn more</a>

                                <div class="sm:flex sm:items-center ">
                                    <button @click="() => { modalState = false }" class="w-full px-4 py-2 mt-2 text-sm font-medium tracking-wide text-gray-700 capitalize transition-colors duration-300 transform border border-gray-200 rounded-md sm:mt-0 sm:w-auto sm:mx-2 dark:text-gray-200 dark:border-gray-700 dark:hover:bg-gray-800 hover:bg-gray-100 focus:outline-none focus:ring focus:ring-gray-300 focus:ring-opacity-40">
                                        Cancel
                                    </button>

                                    <button class="w-full px-4 py-2 mt-2 text-sm font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-md sm:w-auto sm:mt-0 hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40">
                                        Archive
                                    </button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </transition>
        </div>
    </div>
</template>

<script setup>
    import Swal from 'sweetalert2'
    const swalWithTailwind = Swal.mixin({
        customClass: {
            popup: 'bg-white dark:bg-gray-900',
            confirmButton: 'px-8 py-2.5 leading-5 text-white transition-colors duration-300 transform bg-gray-700 rounded-md hover:bg-gray-600 focus:outline-none focus:bg-gray-600',
            cancelButton: 'px-8 py-2.5 leading-5 text-white transition-colors duration-300 transform bg-gray-700 rounded-md hover:bg-gray-600 focus:outline-none focus:bg-gray-600'
        },
        buttonsStyling: false,
    })

    const showUserLogin = ref(false)
    const cookieStatus = useState('cookieState', () => false)
    const modalState = ref(false)

    const accountSave = (event) => {
        event.preventDefault()
        console.log("test")
    }
</script>