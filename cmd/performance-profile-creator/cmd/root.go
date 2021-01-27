/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	// "context"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"k8s.io/utils/pointer"
	"k8s.io/klog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	performancev2 "github.com/openshift-kni/performance-addon-operators/api/v2"
	testutils "github.com/openshift-kni/performance-addon-operators/functests/utils"
	// testclient "github.com/openshift-kni/performance-addon-operators/functests/utils/client"
	// homedir "github.com/mitchellh/go-homedir"
	// "github.com/spf13/viper"
)

var (
	powerConsumptionMode, outputFilePath string
	reservedCPUCount int
	splitCPUsAcrossNUMA, disableHT, rtKernel, userLevelNetworking bool
)
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "performance-profile-creator",
	Short: "A tool that automates creation of Performance Profiles",
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// cobra.OnInitialize(initArgs)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().IntVar(&reservedCPUCount, "reserved-cpu-count", 1, "Number of reserved CPUs")
	rootCmd.PersistentFlags().BoolVar(&splitCPUsAcrossNUMA, "split-cpus-across-numa", true, "Split the CPUs across NUMA nodes")
  rootCmd.PersistentFlags().BoolVar(&disableHT, "disable-ht", false, "Disable Hyperthreading")
  rootCmd.PersistentFlags().BoolVar(&rtKernel, "rt-kernel", true, "Enable Real Time Kernel")
  rootCmd.PersistentFlags().BoolVar(&userLevelNetworking, "user-level-networking", false, "Run with User level Networking(DPDK) enabled")
  rootCmd.PersistentFlags().StringVar(&powerConsumptionMode, "power-consumption-mode", "cstate", "The power consumption mode")
	rootCmd.PersistentFlags().StringVar(&outputFilePath, "output-file-path", "/root/go/src/github.com/openshift-kni/performance-addon-operators/build/_output", "Performance Profile file output path")

	// TODO: Check if power consumtion either CSTATE NO_CSTATE IDLE_POLL

	// Sample Performance Profile
 createSampleProfile(outputFilePath)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createSampleProfile(outputFilePath string){
	newRole := "worker-node"
	newLabel := fmt.Sprintf("%s/%s", testutils.LabelRole, newRole)

	reserved := performancev2.CPUSet("0")
	isolated := performancev2.CPUSet("1-3")

	sampleProfile := &performancev2.PerformanceProfile{
		TypeMeta: metav1.TypeMeta{
			Kind:       "PerformanceProfile",
			APIVersion: performancev2.GroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "sample-profile",
		},
		Spec: performancev2.PerformanceProfileSpec{
			CPU: &performancev2.CPU{
				Reserved: &reserved,
				Isolated: &isolated,
			},
			NodeSelector: map[string]string{newLabel: ""},
			RealTimeKernel: &performancev2.RealTimeKernel{
				Enabled: pointer.BoolPtr(true),
			},
			AdditionalKernelArgs: []string{
				"NEW_ARGUMENT",
			},
			NUMA: &performancev2.NUMA{
				TopologyPolicy: pointer.StringPtr("restricted"),
			},
		},
	}

	 var performanceProfileData []byte
	 var err error

	if performanceProfileData, err = json.Marshal(&sampleProfile); err != nil {
		klog.Errorf("Unable to Marshal sample performance profile: %v", err)
    }

	klog.Infof("The performance profile data is %q", performanceProfileData)

	// testclient.Client.Create(context.TODO(), sampleProfile)
}
//
// // initConfig reads in config file and ENV variables if set.
// func initArgs() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}
//
// 		// Search config in home directory with name ".performance-profile-creator" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".performance-profile-creator")
// 	}
//
// 	viper.AutomaticEnv() // read in environment variables that match
//
// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
