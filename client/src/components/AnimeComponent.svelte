<script lang="ts">
    export let title: string;
    export let mal_id: number;
    export let image: string;
    export let height = "200px";
    export let width = "200px";
</script>

<a href="/{mal_id}" {title}>
    <img
        class="hover:opacity-50 hover:cursor-pointer"
        src={image}
        alt={title}
        {height}
        {width}
    />
</a>
