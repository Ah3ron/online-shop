<script>
  import { onMount } from "svelte";

  let products = [];
  let loading = true;
  let error = null;

  async function fetchProducts() {
    try {
      const response = await fetch("http://localhost:3000/api/products");
      if (!response.ok) {
        throw new Error("Ошибка при загрузке данных");
      }
      products = await response.json();
    } catch (err) {
      error = err.message;
    } finally {
      loading = false;
    }
  }

  onMount(() => {
    fetchProducts();
  });
</script>

<svelte:head>
  <title>Каталог - Магазин Кубиков Рубика</title>
  <meta name="description" content="Каталог кубиков Рубика" />
</svelte:head>

<section class="p-4 mx-full bg-base-100">
  <h1 class="text-3xl font-bold text-center mb-6">Каталог Кубиков Рубика</h1>

  {#if loading}
    <p class="text-center">Загрузка...</p>
  {:else if error}
    <p class="text-error-500 text-center">Ошибка: {error}</p>
  {:else if products.length === 0}
    <p class="text-center">Нет доступных товаров.</p>
  {:else}
    <div
      class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 shadow-2xl"
    >
      {#each products as product}
        <div class="card bg-base-200 rounded-lg shadow-lg p-4">
          <h2 class="text-xl font-semibold mb-2">{product.name}</h2>
          <p class="text-lg font-bold mb-4">Цена: {product.price} ₽</p>
          <button class="btn btn-primary w-full">Добавить в корзину</button>
        </div>
      {/each}
    </div>
  {/if}
</section>
