<script lang="ts">
	import { PUBLIC_PUBLIC_URL } from '$env/static/public';

	export let id;
</script>

<p class="mt-4 text-center text-lg font-semibold dark:text-gray-200">
	This is how the form will look like to the users.
</p>
<p class="mt-2 text-center font-bold text-red-600">
	WARNING: This is a live view of the form. Any submissions will be recorded.
</p>

<iframe
	src={(PUBLIC_PUBLIC_URL ?? 'http://localhost:8889') + '/form/' + id}
	title="Form View"
	class="h-screen w-full"
></iframe>
