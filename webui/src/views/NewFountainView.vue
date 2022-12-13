<script>
export default {
	data: function () {
		return {
			errormsg: null,
			loading: false,
			latitude: null,
			longitude: null,
			fountainStatus: "good",
		}
	},
	methods: {
		createFountain: async function () {
			this.loading = true;
			this.errormsg = null;
			try {
				await this.$axios.post("/fountains/", {
					status: this.fountainStatus,
					latitude: this.latitude,
					longitude: this.longitude,
				});
				this.$router.push("/");
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		}
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">New Fountain</h1>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div class="mb-3">
			<label for="description" class="form-label">Latitude</label>
			<input type="number" class="form-control" id="latitude" v-model="latitude" placeholder="12.34">
		</div>
		<div class="mb-3">
			<label for="description" class="form-label">Longitude</label>
			<input type="number" class="form-control" id="longitude" v-model="longitude" placeholder="56.78">
		</div>
		<div class="form-check">
			<input class="form-check-input" type="radio" id="statusGood" value="good" v-model="fountainStatus" checked>
			<label class="form-check-label" for="statusGood">Good</label>
		</div>
		<div class="form-check">
			<input class="form-check-input" type="radio" id="statusGood" value="faulty" v-model="fountainStatus">
			<label class="form-check-label" for="statusGood">Faulty</label>
		</div>

		<div>
			<button v-if="!loading" type="button" class="btn btn-primary" @click="createFountain">
				Create fountain
			</button>
			<LoadingSpinner v-if="loading"></LoadingSpinner>
		</div>
	</div>
</template>

<style>
</style>
