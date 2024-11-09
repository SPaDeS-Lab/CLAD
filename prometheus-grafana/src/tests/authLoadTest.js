const { sleep } = require('../utils/helpers');
const GrafanaAPI = require('../monitoring/grafanaApi');

async function performAuthLoadTest(requestsPerMinute, isOffline) {
    const testDuration = 60000;
    const delayBetweenRequests = 60000 / requestsPerMinute;
    const results = [];

    console.log(
        `Starting auth load test: ${requestsPerMinute} requests/minute, ${
            isOffline ? 'offline' : 'online'
        } mode`
    );

    for (let i = 0; i < requestsPerMinute; i++) {
        const startTime = Date.now();

        try {
            if (isOffline) {
                await performOfflineAuth();
            } else {
                await performOnlineAuth();
            }

            results.push({
                latency: Date.now() - startTime,
                success: true,
            });
        } catch (error) {
            results.push({
                latency: Date.now() - startTime,
                success: false,
                error: error.message,
            });
        }

        await sleep(delayBetweenRequests);
    }

    return analyzeResults(results);
}

async function performOfflineAuth() {
    // Implement offline authentication logic
}

async function performOnlineAuth() {
    // Implement online authentication logic
}

function analyzeResults(results) {
    const latencies = results.map((r) => r.latency);
    const successRate =
        (results.filter((r) => r.success).length / results.length) * 100;

    return {
        avgLatency: latencies.reduce((a, b) => a + b, 0) / latencies.length,
        maxLatency: Math.max(...latencies),
        minLatency: Math.min(...latencies),
        successRate,
        totalRequests: results.length,
    };
}

module.exports = {
    performAuthLoadTest,
};
