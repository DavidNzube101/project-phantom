<script>
	import { Card, CardContent } from '$lib/components/ui/card';
</script>

<Card class="border-purple-700/30 bg-purple-900/20 backdrop-blur-sm shadow-sm animate-pulse">
	<CardContent class="p-6">
		<div class="flex items-start gap-4">
			<div class="flex-shrink-0 h-8 w-8 rounded-lg bg-purple-800/50"></div>
			<div class="flex-1 space-y-3">
				<div class="h-4 bg-purple-800/50 rounded w-3/4"></div>
				<div class="h-6 bg-purple-800/50 rounded w-full"></div>
				<div class="h-4 bg-purple-800/50 rounded w-5/6"></div>
			</div>
		</div>
	</CardContent>
</Card>
