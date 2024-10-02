<script>
  const welcome_url = "https://placehold.co/600x400/png";
  const welcome_fallback_url = "https://placehold.co/600x400/png";
</script>

<svelte:head>
  <title>Главная - Магазин Кубиков Рубика</title>
  <meta name="description" content="Онлайн магазин кубиков Рубика" />
</svelte:head>

<section class="flex h-max bg-base-100">
  <div class="flex-1 flex flex-col justify-center items-start p-10">
    <h1 class="text-4xl font-bold mb-4">
      Добро пожаловать в наш онлайн магазин кубиков Рубика
    </h1>
    <p class="text-lg mb-6">
      Откройте для себя широкий ассортимент кубиков Рубика и аксессуаров!
    </p>
    <a href="/products" class="btn btn-primary">Посмотреть каталог</a>
  </div>

  <div class="flex-1 flex justify-center items-center">
    <picture>
      <source srcset={welcome_url} type="image/webp" />
      <img src={welcome_fallback_url} alt="Кубик Рубика" class="w-3/4 h-auto" />
    </picture>
  </div>
</section>
