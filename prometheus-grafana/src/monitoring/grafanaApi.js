class GrafanaAPI {
    constructor(config) {
        this.baseUrl = config.baseUrl;
        this.apiKey = config.apiKey;
    }

    async collectMetrics({ nodes, metrics, interval }) {
        const results = {};

        for (const nodeId of nodes) {
            results[nodeId] = await this.queryMetrics({
                nodeId,
                metrics,
                interval,
            });
        }

        return results;
    }

    async queryMetrics({ nodeId, metrics, interval }) {
        // Implementation for querying Grafana API
        const queries = metrics.map((metric) => ({
            target: `${metric}.${nodeId}`,
            interval,
        }));

        const response = await fetch(`${this.baseUrl}/api/ds/query`, {
            method: 'POST',
            headers: {
                Authorization: `Bearer ${this.apiKey}`,
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ queries }),
        });

        return response.json();
    }
}

module.exports = GrafanaAPI;
