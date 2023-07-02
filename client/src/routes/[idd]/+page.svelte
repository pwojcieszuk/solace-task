<script lang="ts">
    import AnimeComponent from "../../components/AnimeComponent.svelte";
    import { favorites } from "../../stores";
    import type { PageData } from "./$types";

    export let data: PageData;
</script>

<AnimeComponent
    title={data.anime.title}
    mal_id={data.anime.mal_id}
    image={data.anime.images.webp.image_url}
/>
<div class="mt-10">
    <button
        class="bg-gray-300 rounded p-4 mt-4"
        on:click={() =>
            favorites.add(
                data.anime.mal_id,
                data.anime.title,
                data.anime.images.webp.image_url,
            )}
    >
        Add to favorites
    </button>

    <a class="bg-gray-300 rounded p-4" href="/">Go back to list</a>
</div>
