<script lang="ts">
	import '../app.css';
	import { goto } from '$app/navigation';
	import { Button, DarkMode, Navbar, NavBrand, Alert } from 'flowbite-svelte';
	import { onMount, onDestroy } from 'svelte';
	import { activeAlert, type AlertType } from '$lib/alert.ts';

	let alert: AlertType | undefined = undefined;
	let unsubscribe = () => {};
	onMount(() => {
		unsubscribe = activeAlert.subscribe((value: AlertType | undefined) => {
			alert = value;
		});
	});

	onDestroy(() => {
		unsubscribe();
	});
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

{#if alert}
	<Alert
		color={alert.type === 'success' ? 'green' : alert.type === 'warn' ? 'yellow' : 'red'}
		class="fixed right-[50%] top-0 m-4 translate-x-[50%]"
		on:click={() => {
			activeAlert.set(undefined);
		}}
		dismissable
	>
		{alert.message}
	</Alert>
{/if}
