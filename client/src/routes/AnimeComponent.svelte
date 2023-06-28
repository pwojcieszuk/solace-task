<script lang="ts">
    export let title: string;
    export let mal_id: number;
    export let image: string;
</script>

<a href="/{mal_id}">
    <img
        class="hover:opacity-50 hover:cursor-pointer mx-auto"
        src={image}
        alt={title}
        height="200px"
        width="200px"
    />
</a>
<h1>{title}</h1>
