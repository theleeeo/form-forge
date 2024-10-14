<script>
	import '../app.css';
	import { goto } from '$app/navigation';

	import { Button, DarkMode, Navbar, NavBrand } from 'flowbite-svelte';
</script>

<!-- <nav>
	<h1 class="m-4 text-primary-600">Welcome to FormForge</h1>
	<div class="centered-group"></div>
</nav> -->

<Navbar>
	<NavBrand href="/">
		<!-- <img src="/images/flowbite-svelte-icon-logo.svg" class="me-3 h-6 sm:h-9" alt="Flowbite Logo" /> -->
		<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
			>Welcome to Form-Forge!</span
		>
	</NavBrand>
	<!-- <NavHamburger />
	<NavUl>
		<NavLi href="/">Home</NavLi>
		<NavLi href="/about">About</NavLi>
		<NavLi href="/docs/components/navbar">Navbar</NavLi>
		<NavLi href="/pricing">Pricing</NavLi>
		<NavLi href="/contact">Contact</NavLi>
	</NavUl> -->
	<div class="centered-group">
		<DarkMode></DarkMode>
		<Button size="xl" class="m-1 border border-black px-10" on:click={() => goto('/create')}
			>Create New</Button
		>
	</div>
</Navbar>

<main>
	<slot />
</main>

<style>
	main {
		padding: 1em;
		width: 100%;

		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.centered-group {
		display: flex;
		align-items: center;

		gap: 1em;
	}
</style>
