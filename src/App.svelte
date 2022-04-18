<svelte:options tag="my-app" />

<script>
	export let name;
</script>

<main>
	<h1>Hello {name}!</h1>
	<p>
		Visit the
		<a href="https://svelte.dev/tutorial">Svelte tutorial</a>
		to learn how to build Svelte apps.
	</p>
	<form action="/input" method="POST">
		<label for="fname">First name:</label><br />
		<input type="text" id="fname" name="fname" /><br />
		<label for="lname">Last name:</label><br />
		<input type="text" id="lname" name="lname" />

		<label for="cars">Choose a car:</label>
		<select name="cars" id="cars">
			<option value="volvo">Volvo</option>
			<option value="saab">Saab</option>
			<option value="mercedes">Mercedes</option>
			<option value="audi">Audi</option>
		</select>
		<input type="submit" />
	</form>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}
	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}
	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
