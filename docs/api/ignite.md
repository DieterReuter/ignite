

# v1alpha1
`import "/go/src/github.com/weaveworks/ignite/pkg/apis/ignite/v1alpha1"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
+k8s:deepcopy-gen=package
+k8s:defaulter-gen=TypeMeta
+k8s:openapi-gen=true
+k8s:conversion-gen=github.com/weaveworks/ignite/pkg/apis/ignite




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [Variables](#pkg-variables)
* [func SetDefaults_OCIImageClaim(obj *OCIImageClaim)](#SetDefaults_OCIImageClaim)
* [func SetDefaults_PoolSpec(obj *PoolSpec)](#SetDefaults_PoolSpec)
* [func SetDefaults_VMKernelSpec(obj *VMKernelSpec)](#SetDefaults_VMKernelSpec)
* [func SetDefaults_VMNetworkSpec(obj *VMNetworkSpec)](#SetDefaults_VMNetworkSpec)
* [func SetDefaults_VMSpec(obj *VMSpec)](#SetDefaults_VMSpec)
* [func SetDefaults_VMStatus(obj *VMStatus)](#SetDefaults_VMStatus)
* [type FileMapping](#FileMapping)
* [type Image](#Image)
* [type ImageSourceType](#ImageSourceType)
* [type ImageSpec](#ImageSpec)
* [type ImageStatus](#ImageStatus)
* [type Kernel](#Kernel)
* [type KernelSpec](#KernelSpec)
* [type KernelStatus](#KernelStatus)
* [type NetworkMode](#NetworkMode)
  * [func (nm NetworkMode) String() string](#NetworkMode.String)
* [type OCIImageClaim](#OCIImageClaim)
* [type OCIImageSource](#OCIImageSource)
* [type Pool](#Pool)
* [type PoolDevice](#PoolDevice)
* [type PoolDeviceType](#PoolDeviceType)
* [type PoolSpec](#PoolSpec)
* [type PoolStatus](#PoolStatus)
* [type SSH](#SSH)
  * [func (s *SSH) MarshalJSON() ([]byte, error)](#SSH.MarshalJSON)
  * [func (s *SSH) UnmarshalJSON(b []byte) error](#SSH.UnmarshalJSON)
* [type VM](#VM)
* [type VMImageSpec](#VMImageSpec)
* [type VMKernelSpec](#VMKernelSpec)
* [type VMNetworkSpec](#VMNetworkSpec)
* [type VMSpec](#VMSpec)
* [type VMState](#VMState)
* [type VMStatus](#VMStatus)


#### <a name="pkg-files">Package files</a>
[defaults.go](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/defaults.go) [doc.go](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/doc.go) [json.go](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/json.go) [register.go](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/register.go) [types.go](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    KindImage  meta.Kind = "Image"
    KindKernel meta.Kind = "Kernel"
    KindVM     meta.Kind = "VM"
)
```
``` go
const (
    // GroupName is the group name use in this package
    GroupName = "ignite.weave.works"
)
```

## <a name="pkg-variables">Variables</a>
``` go
var (
    // SchemeBuilder the schema builder
    SchemeBuilder = runtime.NewSchemeBuilder(
        addKnownTypes,
        addDefaultingFuncs,
    )

    AddToScheme = localSchemeBuilder.AddToScheme
)
```
``` go
var SchemeGroupVersion = schema.GroupVersion{
    Group:   GroupName,
    Version: "v1alpha1",
}
```
SchemeGroupVersion is group version used to register these objects



## <a name="SetDefaults_OCIImageClaim">func</a> [SetDefaults_OCIImageClaim](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/defaults.go?s=263:313#L13)
``` go
func SetDefaults_OCIImageClaim(obj *OCIImageClaim)
```


## <a name="SetDefaults_PoolSpec">func</a> [SetDefaults_PoolSpec](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/defaults.go?s=353:393#L17)
``` go
func SetDefaults_PoolSpec(obj *PoolSpec)
```


## <a name="SetDefaults_VMKernelSpec">func</a> [SetDefaults_VMKernelSpec](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/defaults.go?s=1235:1283#L53)
``` go
func SetDefaults_VMKernelSpec(obj *VMKernelSpec)
```


## <a name="SetDefaults_VMNetworkSpec">func</a> [SetDefaults_VMNetworkSpec](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/defaults.go?s=1520:1570#L64)
``` go
func SetDefaults_VMNetworkSpec(obj *VMNetworkSpec)
```


## <a name="SetDefaults_VMSpec">func</a> [SetDefaults_VMSpec](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/defaults.go?s=919:955#L39)
``` go
func SetDefaults_VMSpec(obj *VMSpec)
```


## <a name="SetDefaults_VMStatus">func</a> [SetDefaults_VMStatus](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/defaults.go?s=1641:1681#L70)
``` go
func SetDefaults_VMStatus(obj *VMStatus)
```



## <a name="FileMapping">type</a> [FileMapping](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=7677:7772#L206)
``` go
type FileMapping struct {
    HostPath string `json:"hostPath"`
    VMPath   string `json:"vmPath"`
}

```
FileMapping defines mappings between files on the host and VM










## <a name="Image">type</a> [Image](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=342:806#L17)
``` go
type Image struct {
    meta.TypeMeta `json:",inline"`
    // meta.ObjectMeta is also embedded into the struct, and defines the human-readable name, and the machine-readable ID
    // Name is available at the .metadata.name JSON path
    // ID is available at the .metadata.uid JSON path (the Go type is k8s.io/apimachinery/pkg/types.UID, which is only a typed string)
    meta.ObjectMeta `json:"metadata"`

    Spec   ImageSpec   `json:"spec"`
    Status ImageStatus `json:"status"`
}

```
Image represents a cached OCI image ready to be used with Ignite
+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object










## <a name="ImageSourceType">type</a> [ImageSourceType](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=995:1022#L34)
``` go
type ImageSourceType string
```
ImageSourceType is an enum of different supported Image Source Types


``` go
const (
    // ImageSourceTypeDocker defines that the image is imported from Docker
    ImageSourceTypeDocker ImageSourceType = "Docker"
)
```









## <a name="ImageSpec">type</a> [ImageSpec](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=854:921#L29)
``` go
type ImageSpec struct {
    OCIClaim OCIImageClaim `json:"ociClaim"`
}

```
ImageSpec declares what the image contains










## <a name="ImageStatus">type</a> [ImageStatus](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=2320:2469#L68)
``` go
type ImageStatus struct {
    // OCISource contains the information about how this OCI image was imported
    OCISource OCIImageSource `json:"ociSource"`
}

```
ImageStatus defines the status of the image










## <a name="Kernel">type</a> [Kernel](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=4848:5315#L132)
``` go
type Kernel struct {
    meta.TypeMeta `json:",inline"`
    // meta.ObjectMeta is also embedded into the struct, and defines the human-readable name, and the machine-readable ID
    // Name is available at the .metadata.name JSON path
    // ID is available at the .metadata.uid JSON path (the Go type is k8s.io/apimachinery/pkg/types.UID, which is only a typed string)
    meta.ObjectMeta `json:"metadata"`

    Spec   KernelSpec   `json:"spec"`
    Status KernelStatus `json:"status"`
}

```
Kernel is a serializable object that caches information about imported kernels
This file is stored in /var/lib/firecracker/kernels/{oci-image-digest}/metadata.json
+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object










## <a name="KernelSpec">type</a> [KernelSpec](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=5368:5541#L144)
``` go
type KernelSpec struct {
    OCIClaim OCIImageClaim `json:"ociClaim"`
}

```
KernelSpec describes the properties of a kernel










## <a name="KernelStatus">type</a> [KernelStatus](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=5592:5708#L151)
``` go
type KernelStatus struct {
    Version   string         `json:"version"`
    OCISource OCIImageSource `json:"ociSource"`
}

```
KernelStatus describes the status of a kernel










## <a name="NetworkMode">type</a> [NetworkMode](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=8126:8149#L221)
``` go
type NetworkMode string
```
NetworkMode defines different states a VM can be in


``` go
const (
    // NetworkModeCNI specifies the network mode where CNI is used
    NetworkModeCNI NetworkMode = "cni"
    // NetworkModeDockerBridge specifies the default docker bridge network is used
    NetworkModeDockerBridge NetworkMode = "docker-bridge"
)
```









### <a name="NetworkMode.String">func</a> (NetworkMode) [String](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=8189:8226#L225)
``` go
func (nm NetworkMode) String() string
```



## <a name="OCIImageClaim">type</a> [OCIImageClaim](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=1218:1636#L42)
``` go
type OCIImageClaim struct {
    // Type defines how the image should be imported
    Type ImageSourceType `json:"type"`
    // Ref defines the reference to use when talking to the backend.
    // This is most commonly the image name, followed by a tag.
    // Other supported ways are $registry/$user/$image@sha256:$digest
    // This ref is also used as ObjectMeta.Name for kinds Images and Kernels
    Ref meta.OCIImageRef `json:"ref"`
}

```
OCIImageClaim defines a claim for importing an OCI image










## <a name="OCIImageSource">type</a> [OCIImageSource](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=1743:2271#L54)
``` go
type OCIImageSource struct {
    // ID defines the source's ID (e.g. the Docker image ID)
    ID string `json:"id"`
    // Size defines the size of the source in bytes
    Size meta.Size `json:"size"`
    // RepoDigests defines the image name as it was when pulled
    // from a repository, and the digest of the image
    // The format is $registry/$user/$image@sha256:$digest
    // This field is unpopulated if the image used as the source
    // has never been pushed to or pulled from a registry
    RepoDigests []string `json:"repoDigests,omitempty"`
}

```
OCIImageSource specifies how the OCI image was imported.
It is the status variant of OCIImageClaim










## <a name="Pool">type</a> [Pool](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=2744:2924#L77)
``` go
type Pool struct {
    meta.TypeMeta `json:",inline"`

    Spec   PoolSpec   `json:"spec"`
    Status PoolStatus `json:"status"`
}

```
Pool defines device mapper pool database
This file is managed by the snapshotter part of Ignite, and the file (existing as a singleton)
is present at /var/lib/firecracker/snapshotter/pool.json
+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object










## <a name="PoolDevice">type</a> [PoolDevice](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=4215:4605#L119)
``` go
type PoolDevice struct {
    Size   meta.Size `json:"size"`
    Parent meta.DMID `json:"parent"`
    // Type specifies the type of the contents of the device
    Type PoolDeviceType `json:"type"`
    // MetadataPath points to the JSON/YAML file with metadata about this device
    // This is most often of the format /var/lib/firecracker/{type}/{id}/metadata.json
    MetadataPath string `json:"metadataPath"`
}

```
PoolDevice defines one device in the pool










## <a name="PoolDeviceType">type</a> [PoolDeviceType](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=3944:3970#L109)
``` go
type PoolDeviceType string
```

``` go
const (
    PoolDeviceTypeImage  PoolDeviceType = "Image"
    PoolDeviceTypeResize PoolDeviceType = "Resize"
    PoolDeviceTypeKernel PoolDeviceType = "Kernel"
    PoolDeviceTypeVM     PoolDeviceType = "VM"
)
```









## <a name="PoolSpec">type</a> [PoolSpec](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=2971:3684#L87)
``` go
type PoolSpec struct {
    // MetadataSize specifies the size of the pool's metadata
    MetadataSize meta.Size `json:"metadataSize"`
    // DataSize specifies the size of the pool's data
    DataSize meta.Size `json:"dataSize"`
    // AllocationSize specifies the smallest size that can be allocated at a time
    AllocationSize meta.Size `json:"allocationSize"`
    // MetadataPath points to the file where device mapper stores all metadata information
    // Defaults to constants.SNAPSHOTTER_METADATA_PATH
    MetadataPath string `json:"metadataPath"`
    // DataPath points to the backing physical device or sparse file (to be loop mounted) for the pool
    // Defaults to constants.SNAPSHOTTER_DATA_PATH
    DataPath string `json:"dataPath"`
}

```
PoolSpec defines the Pool's specification










## <a name="PoolStatus">type</a> [PoolStatus](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=3734:3942#L103)
``` go
type PoolStatus struct {
    // The Devices array needs to contain pointers to accommodate "holes" in the mapping
    // Where devices have been deleted, the pointer is nil
    Devices []*PoolDevice `json:"devices"`
}

```
PoolStatus defines the Pool's current status










## <a name="SSH">type</a> [SSH](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=7992:8069#L215)
``` go
type SSH struct {
    Generate  bool   `json:"-"`
    PublicKey string `json:"-"`
}

```
SSH specifies different ways to connect via SSH to the VM
SSH uses a custom marshaller/unmarshaller. If generate is true,
it marshals to true (a JSON bool). If PublicKey is set, it marshals
to that string.










### <a name="SSH.MarshalJSON">func</a> (\*SSH) [MarshalJSON](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/json.go?s=117:160#L9)
``` go
func (s *SSH) MarshalJSON() ([]byte, error)
```



### <a name="SSH.UnmarshalJSON">func</a> (\*SSH) [UnmarshalJSON](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/json.go?s=308:351#L21)
``` go
func (s *SSH) UnmarshalJSON(b []byte) error
```



## <a name="VM">type</a> [VM](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=5910:6365#L159)
``` go
type VM struct {
    meta.TypeMeta `json:",inline"`
    // meta.ObjectMeta is also embedded into the struct, and defines the human-readable name, and the machine-readable ID
    // Name is available at the .metadata.name JSON path
    // ID is available at the .metadata.uid JSON path (the Go type is k8s.io/apimachinery/pkg/types.UID, which is only a typed string)
    meta.ObjectMeta `json:"metadata"`

    Spec   VMSpec   `json:"spec"`
    Status VMStatus `json:"status"`
}

```
VM represents a virtual machine run by Firecracker
These files are stored in /var/lib/firecracker/vm/{vm-id}/metadata.json
+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object










## <a name="VMImageSpec">type</a> [VMImageSpec](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=7298:7367#L191)
``` go
type VMImageSpec struct {
    OCIClaim OCIImageClaim `json:"ociClaim"`
}

```









## <a name="VMKernelSpec">type</a> [VMKernelSpec](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=7369:7490#L195)
``` go
type VMKernelSpec struct {
    OCIClaim OCIImageClaim `json:"ociClaim"`
    CmdLine  string        `json:"cmdLine,omitempty"`
}

```









## <a name="VMNetworkSpec">type</a> [VMNetworkSpec](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=7492:7610#L200)
``` go
type VMNetworkSpec struct {
    Mode  NetworkMode       `json:"mode"`
    Ports meta.PortMappings `json:"ports,omitempty"`
}

```









## <a name="VMSpec">type</a> [VMSpec](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=6413:7296#L171)
``` go
type VMSpec struct {
    Image    VMImageSpec   `json:"image"`
    Kernel   VMKernelSpec  `json:"kernel"`
    CPUs     uint64        `json:"cpus"`
    Memory   meta.Size     `json:"memory"`
    DiskSize meta.Size     `json:"diskSize"`
    Network  VMNetworkSpec `json:"network"`

    // This will be done at either "ignite start" or "ignite create" time
    // TODO: We might revisit this later
    CopyFiles []FileMapping `json:"copyFiles,omitempty"`
    // SSH specifies how the SSH setup should be done
    // nil here means "don't do anything special"
    // If SSH.Generate is set, Ignite will generate a new SSH key and copy it in to authorized_keys in the VM
    // Specifying a path in SSH.Generate means "use this public key"
    // If SSH.PublicKey is set, this struct will marshal as a string using that path
    // If SSH.Generate is set, this struct will marshal as a bool => true
    SSH *SSH `json:"ssh,omitempty"`
}

```
VMSpec describes the configuration of a VM










## <a name="VMState">type</a> [VMState](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=8623:8642#L238)
``` go
type VMState string
```
VMState defines different states a VM can be in


``` go
const (
    VMStateCreated VMState = "Created"
    VMStateRunning VMState = "Running"
    VMStateStopped VMState = "Stopped"
)
```









## <a name="VMStatus">type</a> [VMStatus](https://github.com/weaveworks/ignite/tree/master/pkg/apis/ignite/v1alpha1/types.go?s=8802:9023#L247)
``` go
type VMStatus struct {
    State       VMState          `json:"state"`
    IPAddresses meta.IPAddresses `json:"ipAddresses,omitempty"`
    Image       OCIImageSource   `json:"image"`
    Kernel      OCIImageSource   `json:"kernel"`
}

```
VMStatus defines the status of a VM














- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
