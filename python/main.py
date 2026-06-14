"""
main.py — mirrors main.go: a registry that runs all modules or one by name.
Usage: python main.py [module_name]
"""
import sys

# Import every module
from statistics.module import Module as Statistics
from probability.module import Module as Probability
from ab_testing.module import Module as ABTesting
from linear_regression.module import Module as LinearRegression
from logistic_regression.module import Module as LogisticRegression
from decision_trees.module import Module as DecisionTrees
from knn.module import Module as KNN
from naive_bayes.module import Module as NaiveBayes
from kmeans.module import Module as KMeans
from data_engineering.module import Module as DataEngineering
from analytics.module import Module as Analytics
from pii_detection.module import Module as PIIDetection
from risk_scoring.module import Module as RiskScoring
from agent_monitoring.module import Module as AgentMonitoring
from hallucination_detection.module import Module as HallucinationDetection


class Registry:
    def __init__(self):
        self._modules = []

    def register(self, module):
        self._modules.append(module)

    def run_all(self):
        for m in self._modules:
            print(f"\n=== {m.name()} ===")
            m.run()

    def run_one(self, name):
        for m in self._modules:
            if m.name() == name:
                print(f"\n=== {m.name()} ===")
                m.run()
                return
        print(f"Module '{name}' not found")
        sys.exit(1)


if __name__ == "__main__":
    reg = Registry()
    reg.register(Statistics())
    reg.register(Probability())
    reg.register(ABTesting())
    reg.register(LinearRegression())
    reg.register(LogisticRegression())
    reg.register(DecisionTrees())
    reg.register(KNN())
    reg.register(NaiveBayes())
    reg.register(KMeans())
    reg.register(DataEngineering())
    reg.register(Analytics())
    reg.register(PIIDetection())
    reg.register(RiskScoring())
    reg.register(AgentMonitoring())
    reg.register(HallucinationDetection())

    if len(sys.argv) > 1:
        reg.run_one(sys.argv[1])
    else:
        reg.run_all()
