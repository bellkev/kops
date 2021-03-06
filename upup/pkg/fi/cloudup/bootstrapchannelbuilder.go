/*
Copyright 2016 The Kubernetes Authors.

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

package cloudup

import (
	"fmt"

	channelsapi "k8s.io/kops/channels/pkg/api"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/fitasks"
	"k8s.io/kops/upup/pkg/fi/utils"
)

type BootstrapChannelBuilder struct {
	cluster *kops.Cluster
}

var _ fi.ModelBuilder = &BootstrapChannelBuilder{}

func (b *BootstrapChannelBuilder) Build(c *fi.ModelBuilderContext) error {
	addons, manifests := b.buildManifest()
	addonsYAML, err := utils.YamlMarshal(addons)
	if err != nil {
		return fmt.Errorf("error serializing addons yaml: %v", err)
	}

	name := b.cluster.ObjectMeta.Name + "-addons-bootstrap"

	tasks := c.Tasks

	tasks[name] = &fitasks.ManagedFile{
		Name:     fi.String(name),
		Location: fi.String("addons/bootstrap-channel.yaml"),
		Contents: fi.WrapResource(fi.NewBytesResource(addonsYAML)),
	}

	for key, resource := range manifests {
		name := b.cluster.ObjectMeta.Name + "-addons-" + key
		tasks[name] = &fitasks.ManagedFile{
			Name:     fi.String(name),
			Location: fi.String(resource),
			Contents: &fi.ResourceHolder{Name: resource},
		}
	}

	return nil
}

func (b *BootstrapChannelBuilder) buildManifest() (*channelsapi.Addons, map[string]string) {
	manifests := make(map[string]string)

	addons := &channelsapi.Addons{}
	addons.Kind = "Addons"
	addons.ObjectMeta.Name = "bootstrap"

	{
		key := "core.addons.k8s.io"
		version := "1.4.0"

		location := key + "/v" + version + ".yaml"

		addons.Spec.Addons = append(addons.Spec.Addons, &channelsapi.AddonSpec{
			Name:     fi.String(key),
			Version:  fi.String(version),
			Selector: map[string]string{"k8s-addon": key},
			Manifest: fi.String(location),
		})
		manifests[key] = "addons/" + location
	}

	{
		key := "kube-dns.addons.k8s.io"
		version := "1.5.0"

		location := key + "/v" + version + ".yaml"

		addons.Spec.Addons = append(addons.Spec.Addons, &channelsapi.AddonSpec{
			Name:     fi.String(key),
			Version:  fi.String(version),
			Selector: map[string]string{"k8s-addon": key},
			Manifest: fi.String(location),
		})
		manifests[key] = "addons/" + location
	}

	{
		key := "limit-range.addons.k8s.io"
		version := "1.5.0"

		location := key + "/v" + version + ".yaml"

		addons.Spec.Addons = append(addons.Spec.Addons, &channelsapi.AddonSpec{
			Name:     fi.String(key),
			Version:  fi.String(version),
			Selector: map[string]string{"k8s-addon": key},
			Manifest: fi.String(location),
		})
		manifests[key] = "addons/" + location
	}

	{
		key := "dns-controller.addons.k8s.io"
		version := "1.4.1"

		location := key + "/v" + version + ".yaml"

		addons.Spec.Addons = append(addons.Spec.Addons, &channelsapi.AddonSpec{
			Name:     fi.String(key),
			Version:  fi.String(version),
			Selector: map[string]string{"k8s-addon": key},
			Manifest: fi.String(location),
		})
		manifests[key] = "addons/" + location
	}

	{
		key := "storage-aws.addons.k8s.io"
		version := "1.5.0"

		location := key + "/v" + version + ".yaml"

		addons.Spec.Addons = append(addons.Spec.Addons, &channelsapi.AddonSpec{
			Name:     fi.String(key),
			Version:  fi.String(version),
			Selector: map[string]string{"k8s-addon": key},
			Manifest: fi.String(location),
		})
		manifests[key] = "addons/" + location
	}

	if b.cluster.Spec.Networking.Kopeio != nil {
		key := "networking.kope.io"
		version := "1.0.20161116"

		// TODO: Create configuration object for cni providers (maybe create it but orphan it)?
		location := key + "/v" + version + ".yaml"

		addons.Spec.Addons = append(addons.Spec.Addons, &channelsapi.AddonSpec{
			Name:     fi.String(key),
			Version:  fi.String(version),
			Selector: map[string]string{"k8s-addon": key},
			Manifest: fi.String(location),
		})

		manifests[key] = "addons/" + location
	}

	if b.cluster.Spec.Networking.Weave != nil {
		key := "networking.weave"
		version := "1.8.2"

		// TODO: Create configuration object for cni providers (maybe create it but orphan it)?
		location := key + "/v" + version + ".yaml"

		addons.Spec.Addons = append(addons.Spec.Addons, &channelsapi.AddonSpec{
			Name:     fi.String(key),
			Version:  fi.String(version),
			Selector: map[string]string{"k8s-addon": key},
			Manifest: fi.String(location),
		})

		manifests[key] = "addons/" + location
	}

	if b.cluster.Spec.Networking.Calico != nil {
		key := "networking.projectcalico.org"
		version := "2.0"

		// TODO: Create configuration object for cni providers (maybe create it but orphan it)?
		location := key + "/v" + version + ".yaml"

		addons.Spec.Addons = append(addons.Spec.Addons, &channelsapi.AddonSpec{
			Name:     fi.String(key),
			Version:  fi.String(version),
			Selector: map[string]string{"k8s-addon": key},
			Manifest: fi.String(location),
		})

		manifests[key] = "addons/" + location
	}

	return addons, manifests
}
