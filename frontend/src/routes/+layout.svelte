<script>
  import Header from "./Header.svelte";
  import "../app.css";
</script>

<div class="flex flex-col min-h-screen bg-base-100">
  <Header />

  <main class="flex-1 flex flex-col p-4 max-w-6xl mx-auto box-border">
    <slot />
  </main>

  <footer class="flex flex-col justify-center items-center p-4 bg-base-200">
    <p class="text-center">
      Посетите <a
        href="https://kit.svelte.dev"
        class="font-bold text-info -600 hover:underline">kit.svelte.dev</a
      > чтобы узнать больше о SvelteKit
    </p>
  </footer>
</div>
