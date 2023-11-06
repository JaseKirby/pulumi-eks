package provider

import (
	"fmt"
	"testing"

	"github.com/pulumi/providertest"
)

func TestExamplesUpgrades(t *testing.T) {
	// It would be easy to extract these invocations into a helper `run(t, dir)` but then they
	// wouldn't be recognized as tests anymore.

	t.Run("cluster", func(t *testing.T) {
		runExampleParallel(t, "cluster")
	})

	t.Run("aws-profile", func(t *testing.T) {
		t.Skip("Fails with 'ALT_AWS_PROFILE must be set'")
		test(t, "../examples/aws-profile").Run(t)
	})

	t.Run("aws-profile-role", func(t *testing.T) {
		runExampleParallel(t, "aws-profile-role")
	})

	t.Run("encryption-provider", func(t *testing.T) {
		runExampleParallel(t, "encryption-provider")
	})

	t.Run("cluster-with-serviceiprange", func(t *testing.T) {
		runExampleParallel(t, "cluster-with-serviceiprange")
	})

	t.Run("extra-sg", func(t *testing.T) {
		runExampleParallel(t, "extra-sg")
	})

	t.Run("fargate", func(t *testing.T) {
		runExampleParallel(t, "fargate")
	})

	t.Run("managed-nodegroups", func(t *testing.T) {
		test(t, "../examples/managed-nodegroups").Run(t)
	})

	t.Run("modify-default-eks-sg", func(t *testing.T) {
		t.Skip("upgradetest doesn't understand invoke aws:ec2/getSecurityGroup:getSecurityGroup")
		runExampleParallel(t, "modify-default-eks-sg")
	})

	t.Run("nodegroup", func(t *testing.T) {
		runExampleParallel(t, "nodegroup")
	})

	t.Run("oidc-iam-sa", func(t *testing.T) {
		runExampleParallel(t, "oidc-iam-sa")
	})

	t.Run("scoped-kubeconfigs", func(t *testing.T) {
		runExampleParallel(t, "scoped-kubeconfigs")
	})

	t.Run("storage-classes", func(t *testing.T) {
		runExampleParallel(t, "storage-classes")
	})

	t.Run("subnet-tags", func(t *testing.T) {
		runExampleParallel(t, "subnet-tags")
	})

	t.Run("tags", func(t *testing.T) {
		runExampleParallel(t, "tags")
	})
}

func TestReportUpgradeCoverage(t *testing.T) {
	providertest.ReportUpgradeCoverage(t)
}

func runExampleParallel(t *testing.T, example string, opts ...providertest.Option) {
	t.Parallel()
	test(t, fmt.Sprintf("../examples/%s", example), opts...).Run(t)
}

func test(t *testing.T, dir string, opts ...providertest.Option) *providertest.ProviderTest {
	opts = append(opts,
		providertest.WithProviderName("eks"),

		providertest.WithSkippedUpgradeTestMode(
			providertest.UpgradeTestMode_Quick,
			"Quick mode is only supported for providers written in Go at the moment"),

		providertest.WithBaselineVersion("1.0.4"),
		providertest.WithExtraBaselineDependencies(map[string]string{
			"aws":        "5.42.0",
			"kubernetes": "3.30.2",
		}),

		// This region needs to match the one configured in .github for the CI workflows.
		providertest.WithConfig("aws:region", "us-west-2"),
	)
	return providertest.NewProviderTest(dir, opts...)
}
