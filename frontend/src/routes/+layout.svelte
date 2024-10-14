<script>
	import '../app.css';
	import { goto } from '$app/navigation';

	import { Button, DarkMode, Navbar, NavBrand } from 'flowbite-svelte';
</script>

<Navbar>
	<NavBrand href="/">
		<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
			>Welcome to Form-Forge!</span
		>
	</NavBrand>
	<div class="flex items-center gap-2">
		<DarkMode />
		<Button size="xl" class="m-1 border border-black px-10" on:click={() => goto('/create')}
			>Create New</Button
		>
	</div>
</Navbar>

<main class="flex w-full flex-col items-center justify-center p-4">
	<slot />
</main>
