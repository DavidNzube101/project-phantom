<script>
	import { Cloud, Sun, Thermometer, Droplets, Wind } from 'lucide-svelte';
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';

	let { query } = $props();
	
	// Parse location from query
	function parseLocation(query) {
		const match = query.match(/weather\s+in\s+(.+)/i) || query.match(/(.+)\s+weather/i);
		return match ? match[1].trim() : 'New York';
	}

	const location = parseLocation(query);
	
	// Mock weather data
	const weatherData = {
		location: location,
		temperature: 22,
		condition: 'Partly Cloudy',
		humidity: 65,
		windSpeed: 12,
		forecast: [
			{ day: 'Today', high: 24, low: 18, condition: 'Partly Cloudy' },
			{ day: 'Tomorrow', high: 26, low: 20, condition: 'Sunny' },
			{ day: 'Wednesday', high: 23, low: 17, condition: 'Rainy' }
		]
	};
</script>

<Card class="border-purple-700/30 bg-gradient-to-r from-blue-900/20 to-purple-900/20 backdrop-blur-sm shadow-sm">
	<CardContent class="p-6">
		<div class="flex items-start gap-4">
			<div class="flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br from-blue-500 to-purple-500 flex-shrink-0">
				<Cloud class="h-5 w-5 text-white" />
			</div>
			<div class="flex-1">
				<div class="flex items-center gap-2 mb-4">
					<h3 class="text-lg font-semibold text-white">Weather</h3>
					<Badge variant="secondary" class="bg-blue-500/20 text-blue-400 border-blue-500/30">
						<Thermometer class="h-3 w-3 mr-1" />
						Current
					</Badge>
				</div>
				
				<div class="grid gap-4 md:grid-cols-2">
					<div class="bg-purple-900/30 rounded-lg p-4 border border-purple-700/30">
						<div class="flex items-center justify-between mb-2">
							<span class="text-purple-300 text-sm">{weatherData.location}</span>
							<Sun class="h-5 w-5 text-yellow-400" />
						</div>
						<div class="text-3xl font-bold text-white mb-1">{weatherData.temperature}°C</div>
						<div class="text-purple-300 text-sm">{weatherData.condition}</div>
					</div>
					
					<div class="space-y-3">
						<div class="flex items-center gap-3 text-sm">
							<Droplets class="h-4 w-4 text-blue-400" />
							<span class="text-purple-300">Humidity</span>
							<span class="text-white ml-auto">{weatherData.humidity}%</span>
						</div>
						<div class="flex items-center gap-3 text-sm">
							<Wind class="h-4 w-4 text-gray-400" />
							<span class="text-purple-300">Wind Speed</span>
							<span class="text-white ml-auto">{weatherData.windSpeed} km/h</span>
						</div>
					</div>
				</div>
				
				<div class="mt-4 grid grid-cols-3 gap-2">
					{#each weatherData.forecast as day}
						<div class="bg-purple-900/20 rounded-lg p-3 text-center border border-purple-700/20">
							<div class="text-xs text-purple-300 mb-1">{day.day}</div>
							<div class="text-sm font-semibold text-white">{day.high}°/{day.low}°</div>
							<div class="text-xs text-purple-400 mt-1">{day.condition}</div>
						</div>
					{/each}
				</div>
			</div>
		</div>
	</CardContent>
</Card>
