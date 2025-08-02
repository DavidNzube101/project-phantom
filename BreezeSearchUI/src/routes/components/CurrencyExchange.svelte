<script>
	import { TrendingUp, ArrowRightLeft, DollarSign, RefreshCw } from 'lucide-svelte';
	import { Card, CardContent } from '$lib/components/ui/card';
	import { Badge } from '$lib/components/ui/badge';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';

	let { query } = $props();
	
	// Interactive state
	let fromCurrency = $state('USD');
	let toCurrency = $state('EUR');
	let amount = $state(1);
	let isLoading = $state(false);
	
	// Mock exchange rates
	const exchangeRates = {
		'USD': { 'EUR': 0.85, 'GBP': 0.73, 'JPY': 110.25, 'BTC': 0.000023, 'ETH': 0.00041, 'SOL': 0.0089 },
		'EUR': { 'USD': 1.18, 'GBP': 0.86, 'JPY': 129.50, 'BTC': 0.000027, 'ETH': 0.00048, 'SOL': 0.0105 },
		'GBP': { 'USD': 1.37, 'EUR': 1.16, 'JPY': 150.75, 'BTC': 0.000031, 'ETH': 0.00056, 'SOL': 0.0122 },
		'BTC': { 'USD': 43500, 'EUR': 36900, 'GBP': 32100, 'JPY': 4785000, 'ETH': 17.8, 'SOL': 387.2 },
		'ETH': { 'USD': 2450, 'EUR': 2080, 'GBP': 1790, 'JPY': 269500, 'BTC': 0.056, 'SOL': 21.8 },
		'SOL': { 'USD': 112.5, 'EUR': 95.6, 'GBP': 82.1, 'JPY': 12375, 'BTC': 0.00258, 'ETH': 0.046 }
	};

	const currencies = ['USD', 'EUR', 'GBP', 'JPY', 'BTC', 'ETH', 'SOL'];

	// Parse currency from initial query
	function parseCurrencyQuery(query) {
		const match = query.match(/(\w{3})\s+to\s+(\w{3})/i);
		if (match) {
			return { from: match[1].toUpperCase(), to: match[2].toUpperCase() };
		}
		return { from: 'USD', to: 'EUR' };
	}

	// Initialize from query
	const initialParse = parseCurrencyQuery(query);
	fromCurrency = initialParse.from;
	toCurrency = initialParse.to;

	function getExchangeRate() {
		return exchangeRates[fromCurrency]?.[toCurrency] || 1;
	}

	function getConvertedAmount() {
		return (amount * getExchangeRate()).toFixed(6);
	}

	function swapCurrencies() {
		const temp = fromCurrency;
		fromCurrency = toCurrency;
		toCurrency = temp;
	}

	function refreshRates() {
		isLoading = true;
		// Simulate API call
		setTimeout(() => {
			isLoading = false;
		}, 1000);
	}

	function getCurrencySymbol(currency) {
		const symbols = {
			'USD': '$', 'EUR': '€', 'GBP': '£', 'JPY': '¥',
			'BTC': '₿', 'ETH': 'Ξ', 'SOL': '◎'
		};
		return symbols[currency] || currency;
	}
</script>

<Card class="border-purple-700/30 bg-gradient-to-r from-green-900/20 to-blue-900/20 backdrop-blur-sm shadow-sm">
	<CardContent class="p-6">
		<div class="flex flex-col sm:flex-row items-start sm:items-center gap-4">
			<div class="flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-br from-green-500 to-blue-500 flex-shrink-0">
				<DollarSign class="h-5 w-5 text-white" />
			</div>
			<div class="flex-1 w-full">
				<div class="flex flex-col sm:flex-row items-start sm:items-center gap-2 mb-4">
					<h3 class="text-lg font-semibold text-white">Currency Exchange</h3>
					<Badge variant="secondary" class="bg-green-500/20 text-green-400 border-green-500/30">
						<TrendingUp class="h-3 w-3 mr-1" />
						Live Rates
					</Badge>
					<Button
						variant="ghost"
						size="sm"
						onclick={refreshRates}
						disabled={isLoading}
						class="sm:ml-auto text-purple-300 hover:text-white hover:bg-purple-800/30"
					>
						<RefreshCw class={`h-4 w-4 ${isLoading ? 'animate-spin' : ''}`} />
					</Button>
				</div>
				
				<div class="grid gap-4 md:grid-cols-2 relative">
					<!-- From Currency -->
					<div class="bg-purple-900/30 rounded-lg p-4 border border-purple-700/30">
						<div class="flex items-center gap-2 mb-3">
							<span class="text-purple-300 text-sm">From</span>
							<select 
								bind:value={fromCurrency}
								class="bg-purple-800/50 border border-purple-600/50 rounded px-2 py-1 text-white text-sm focus:outline-none focus:border-purple-500"
							>
								{#each currencies as currency}
									<option value={currency}>{currency}</option>
								{/each}
							</select>
						</div>
						<div class="flex items-center gap-2">
							<span class="text-lg text-purple-300">{getCurrencySymbol(fromCurrency)}</span>
							<Input
								type="number"
								bind:value={amount}
								min="0"
								step="0.01"
								class="bg-transparent border-none text-2xl font-bold text-white p-0 focus:ring-0 focus:border-none"
								placeholder="1.00"
							/>
						</div>
					</div>
					
					<!-- Swap Button -->
					<div class="flex items-center justify-center md:absolute md:left-1/2 md:top-1/2 md:-translate-x-1/2 md:-translate-y-1/2 md:z-10">
						<Button
							variant="outline"
							size="sm"
							onclick={swapCurrencies}
							class="bg-purple-900/50 border-purple-600 text-purple-300 hover:bg-purple-800/50 hover:text-white rounded-full h-10 w-10 p-0"
						>
							<ArrowRightLeft class="h-4 w-4" />
						</Button>
					</div>
					
					<!-- To Currency -->
					<div class="bg-purple-900/30 rounded-lg p-4 border border-purple-700/30">
						<div class="flex items-center gap-2 mb-3">
							<span class="text-purple-300 text-sm">To</span>
							<select 
								bind:value={toCurrency}
								class="bg-purple-800/50 border border-purple-600/50 rounded px-2 py-1 text-white text-sm focus:outline-none focus:border-purple-500"
							>
								{#each currencies as currency}
									<option value={currency}>{currency}</option>
								{/each}
							</select>
						</div>
						<div class="flex items-center gap-2">
							<span class="text-lg text-purple-300">{getCurrencySymbol(toCurrency)}</span>
							<div class="text-2xl font-bold text-green-400">
								{getConvertedAmount()}
							</div>
						</div>
					</div>
				</div>
				
				<div class="mt-4 flex flex-col sm:flex-row items-start sm:items-center justify-between text-sm">
					<div class="flex items-center gap-2 text-purple-300 mb-2 sm:mb-0">
						<ArrowRightLeft class="h-4 w-4" />
						<span>1 {fromCurrency} = {getExchangeRate().toLocaleString()} {toCurrency}</span>
					</div>
					<span class="text-purple-400">Updated 2 minutes ago</span>
				</div>
				
				<!-- Quick Amount Buttons -->
				<div class="mt-4 flex flex-wrap gap-2 items-center">
					<span class="text-sm text-purple-300 self-center">Quick:</span>
					{#each [1, 10, 100, 1000] as quickAmount}
						<Button
							variant="outline"
							size="sm"
							onclick={() => amount = quickAmount}
							class="border-purple-600/50 text-purple-300 hover:bg-purple-800/30 hover:text-white text-xs"
						>
							{quickAmount}
						</Button>
					{/each}
				</div>
			</div>
		</div>
	</CardContent>
</Card>

<style>
	.md\:absolute {
		position: relative;
	}
	
	@media (min-width: 768px) {
		.md\:absolute {
			position: absolute;
		}
	}
</style>
