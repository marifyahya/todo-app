<script setup>
import { ref } from "vue";

const tasks = ref([
  {
    id: 1,
    title: "Research content ideas",
    status: "pending",
    date: "",
    tags: [],
    subtasks: 0,
  },
  {
    id: 2,
    title: "Create a database of guest authors",
    status: "pending",
    date: "",
    tags: [],
    subtasks: 0,
  },
  {
    id: 3,
    title: "Renew driver's license",
    status: "pending",
    date: "22-03-22",
    tags: ["Personal"],
    subtasks: 1,
  },
  {
    id: 4,
    title: "Consult accountant",
    status: "pending",
    date: "",
    tags: ["List 1"],
    subtasks: 3,
  },
  {
    id: 5,
    title: "Print business card",
    status: "pending",
    date: "",
    tags: [],
    subtasks: 0,
  },
]);
</script>

<template>
  <div class="flex min-h-screen bg-white font-sans text-gray-800">
    <!-- SIDEBAR -->
    <aside class="w-72 bg-gray-50 border-r border-gray-200 flex flex-col p-6">
      <div class="flex items-center justify-between mb-8">
        <h1 class="text-2xl font-bold tracking-tight">Menu</h1>
        <button class="text-gray-400 hover:text-gray-600">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-6 w-6"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h16M4 18h16"
            />
          </svg>
        </button>
      </div>

      <!-- Search Bar -->
      <div class="relative mb-8">
        <span
          class="absolute inset-y-0 left-0 pl-3 flex items-center text-gray-400"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
            />
          </svg>
        </span>
        <input
          type="text"
          placeholder="Search"
          class="w-full bg-white border border-gray-200 rounded-lg py-2 pl-10 pr-4 text-sm focus:outline-none focus:ring-2 focus:ring-yellow-400 transition-all"
        />
      </div>

      <!-- Navigation Sections -->
      <div class="space-y-8 overflow-y-auto flex-grow">
        <!-- TASKS -->
        <div>
          <h2
            class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-4"
          >
            Tasks
          </h2>
          <ul class="space-y-1">
            <li
              class="flex items-center justify-between p-2 rounded-lg hover:bg-gray-100 cursor-pointer"
            >
              <div class="flex items-center gap-3">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5 text-gray-400"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M13 5l7 7-7 7M5 5l7 7-7 7"
                  />
                </svg>
                <span class="text-sm font-medium">Upcoming</span>
              </div>
              <span
                class="bg-gray-200 text-gray-600 text-[10px] font-bold px-2 py-0.5 rounded"
                >12</span
              >
            </li>
            <li
              class="flex items-center justify-between p-2 rounded-lg bg-gray-200 cursor-pointer"
            >
              <div class="flex items-center gap-3">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5 text-gray-600"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M4 6h16M4 12h16M4 18h7"
                  />
                </svg>
                <span class="text-sm font-bold">Today</span>
              </div>
              <span
                class="bg-white text-gray-600 text-[10px] font-bold px-2 py-0.5 rounded"
                >5</span
              >
            </li>
            <li
              class="flex items-center gap-3 p-2 rounded-lg hover:bg-gray-100 cursor-pointer text-gray-500"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                />
              </svg>
              <span class="text-sm font-medium">Calendar</span>
            </li>
            <li
              class="flex items-center gap-3 p-2 rounded-lg hover:bg-gray-100 cursor-pointer text-gray-500"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"
                />
              </svg>
              <span class="text-sm font-medium">Sticky Wall</span>
            </li>
          </ul>
        </div>

        <!-- LISTS -->
        <div>
          <h2
            class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-4"
          >
            Lists
          </h2>
          <ul class="space-y-1">
            <li
              class="flex items-center justify-between p-2 rounded-lg hover:bg-gray-100 cursor-pointer"
            >
              <div class="flex items-center gap-3">
                <div class="w-3 h-3 rounded-sm bg-red-400"></div>
                <span class="text-sm font-medium">Personal</span>
              </div>
              <span class="text-gray-400 text-xs font-bold">3</span>
            </li>
            <li
              class="flex items-center justify-between p-2 rounded-lg hover:bg-gray-100 cursor-pointer"
            >
              <div class="flex items-center gap-3">
                <div class="w-3 h-3 rounded-sm bg-cyan-400"></div>
                <span class="text-sm font-medium">Work</span>
              </div>
              <span class="text-gray-400 text-xs font-bold">6</span>
            </li>
            <li
              class="flex items-center justify-between p-2 rounded-lg hover:bg-gray-100 cursor-pointer"
            >
              <div class="flex items-center gap-3">
                <div class="w-3 h-3 rounded-sm bg-yellow-400"></div>
                <span class="text-sm font-medium">List 1</span>
              </div>
              <span class="text-gray-400 text-xs font-bold">3</span>
            </li>
            <li
              class="flex items-center gap-3 p-2 rounded-lg hover:bg-gray-100 cursor-pointer text-gray-400"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M12 4v16m8-8H4"
                />
              </svg>
              <span class="text-sm font-medium">Add New List</span>
            </li>
          </ul>
        </div>

        <!-- TAGS -->
        <div>
          <h2
            class="text-xs font-bold text-gray-400 uppercase tracking-widest mb-4"
          >
            Tags
          </h2>
          <div class="flex flex-wrap gap-2">
            <span
              class="bg-cyan-100 text-cyan-700 px-3 py-1 rounded-md text-xs font-bold cursor-pointer hover:bg-cyan-200 transition-colors"
              >Tag 1</span
            >
            <span
              class="bg-pink-100 text-pink-700 px-3 py-1 rounded-md text-xs font-bold cursor-pointer hover:bg-pink-200 transition-colors"
              >Tag 2</span
            >
            <span
              class="bg-gray-100 text-gray-500 px-3 py-1 rounded-md text-xs font-bold cursor-pointer hover:bg-gray-200 transition-colors"
              >+ Add Tag</span
            >
          </div>
        </div>
      </div>

      <!-- FOOTER -->
      <div class="mt-auto pt-6 border-t border-gray-200 space-y-4">
        <div
          class="flex items-center gap-3 p-2 rounded-lg hover:bg-gray-100 cursor-pointer text-gray-600"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
            />
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
            />
          </svg>
          <span class="text-sm font-medium">Settings</span>
        </div>
        <div
          class="flex items-center gap-3 p-2 rounded-lg hover:bg-gray-100 cursor-pointer text-gray-600"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
            />
          </svg>
          <span class="text-sm font-medium">Sign out</span>
        </div>
      </div>
    </aside>

    <!-- MAIN CONTENT -->
    <main class="flex-grow p-12 bg-white overflow-y-auto">
      <div class="max-w-4xl mx-auto">
        <header class="flex items-center gap-6 mb-12">
          <h2 class="text-6xl font-bold tracking-tight">Today</h2>
          <div
            class="flex items-center justify-center border-2 border-gray-100 rounded-2xl w-20 h-20 text-4xl font-bold text-gray-800"
          >
            5
          </div>
        </header>

        <!-- Add Task Input -->
        <div class="relative group mb-8">
          <div
            class="absolute inset-y-0 left-0 pl-4 flex items-center text-gray-400"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 4v16m8-8H4"
              />
            </svg>
          </div>
          <input
            type="text"
            placeholder="Add New Task"
            class="w-full bg-white border border-gray-100 rounded-xl py-4 pl-12 pr-4 text-lg focus:outline-none focus:border-gray-200 focus:shadow-sm transition-all"
          />
        </div>

        <!-- Task List -->
        <div class="space-y-1">
          <div
            v-for="task in tasks"
            :key="task.id"
            class="group flex items-center justify-between py-4 border-b border-gray-50 hover:bg-gray-50 transition-colors px-4 -mx-4 rounded-xl cursor-pointer"
          >
            <div class="flex flex-col gap-1 w-full">
              <div class="flex items-center gap-4">
                <input
                  type="checkbox"
                  class="w-5 h-5 rounded border-gray-300 text-yellow-400 focus:ring-yellow-400"
                />
                <span class="text-lg font-medium text-gray-700">{{
                  task.title
                }}</span>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-5 w-5 ml-auto text-gray-300 opacity-0 group-hover:opacity-100 transition-opacity"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M9 5l7 7-7 7"
                  />
                </svg>
              </div>

              <!-- Meta Info Row -->
              <div
                v-if="task.date || task.tags.length > 0 || task.subtasks > 0"
                class="flex items-center gap-6 ml-9 mt-1"
              >
                <div
                  v-if="task.date"
                  class="flex items-center gap-1.5 text-gray-500 text-xs font-bold"
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-4 w-4"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                    />
                  </svg>
                  {{ task.date }}
                </div>
                <div
                  v-if="task.subtasks > 0"
                  class="flex items-center gap-1.5 text-gray-500 text-xs font-bold px-2 py-0.5 bg-gray-100 rounded"
                >
                  <span class="text-gray-700">{{ task.subtasks }}</span>
                  Subtasks
                </div>
                <div
                  v-for="tag in task.tags"
                  :key="tag"
                  class="flex items-center gap-1.5 text-gray-700 text-xs font-bold"
                >
                  <div
                    :class="{
                      'bg-red-400': tag === 'Personal',
                      'bg-yellow-400': tag === 'List 1',
                    }"
                    class="w-3 h-3 rounded-sm"
                  ></div>
                  {{ tag }}
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<style>
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap");

body {
  font-family: "Inter", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

/* Custom Checkbox Style */
input[type="checkbox"]:checked {
  background-image: url("data:image/svg+xml,%3csvg viewBox='0 0 16 16' fill='white' xmlns='http://www.w3.org/2000/svg'%3e%3cpath d='M12.207 4.793a1 1 0 010 1.414l-5 5a1 1 0 01-1.414 0l-2-2a1 1 0 011.414-1.414L6.5 9.086l4.293-4.293a1 1 0 011.414 0z'/%3e%3c/svg%3e");
}
</style>
