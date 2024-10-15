<script lang="ts">
	import { PUBLIC_PUBLIC_URL } from '$env/static/public';

	export let id;
</script>

<p class="text-center">This is how the form will look like to the users.</p>

<iframe
	src={(PUBLIC_PUBLIC_URL ?? 'http://localhost:8889') + '/form/' + id}
	title="Form View"
	class="pointer-events-none h-screen w-full"
></iframe>
