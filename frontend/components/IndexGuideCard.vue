<template>
    <div>        
        <div class="max-w-sm overflow-hidden bg-white rounded-lg shadow-md dark:bg-gray-800">
        <img class="object-cover w-full h-64" src="https://images.unsplash.com/photo-1550439062-609e1531270e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=500&q=60" alt="Article">

        <div class="p-4">
            <div class="">
                <h3 class="text-center block mt-2 text-xl font-semibold text-gray-800 transition-colors duration-300 transform dark:text-white">Visit Dashboard</h3>
                <p class="text-center mt-2 text-sm text-gray-600 dark:text-gray-400">Dein Dashboard zeigt dir alle Verfügbaren Infos über den aktuellen Stand deiner Suche</p>

                <div class="flex justify-center mt-4">
                    <button class="flex items-center px-4 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-lg hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80">
                    <svg class="w-5 h-5 mx-1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                        <path fill-rule="evenodd" d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z" clip-rule="evenodd" />
                    </svg>

                    <span class="mx-1">Refresh</span>
                    </button>
                </div>
            </div>
        </div>
        </div>
    </div>
</template>
