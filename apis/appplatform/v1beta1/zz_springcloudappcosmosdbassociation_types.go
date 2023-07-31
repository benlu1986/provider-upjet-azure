/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SpringCloudAppCosmosDBAssociationInitParameters struct {

	// Specifies the API type which should be used when connecting to the CosmosDB Account. Possible values are cassandra, gremlin, mongo, sql or table. Changing this forces a new resource to be created.
	APIType *string `json:"apiType,omitempty" tf:"api_type,omitempty"`

	// Specifies the name of the Cassandra Keyspace which the Spring Cloud App should be associated with. Should only be set when api_type is cassandra.
	CosmosDBCassandraKeySpaceName *string `json:"cosmosdbCassandraKeyspaceName,omitempty" tf:"cosmosdb_cassandra_keyspace_name,omitempty"`

	// Specifies the name of the Gremlin Database which the Spring Cloud App should be associated with. Should only be set when api_type is gremlin.
	CosmosDBGremlinDatabaseName *string `json:"cosmosdbGremlinDatabaseName,omitempty" tf:"cosmosdb_gremlin_database_name,omitempty"`

	// Specifies the name of the Gremlin Graph which the Spring Cloud App should be associated with. Should only be set when api_type is gremlin.
	CosmosDBGremlinGraphName *string `json:"cosmosdbGremlinGraphName,omitempty" tf:"cosmosdb_gremlin_graph_name,omitempty"`

	// Specifies the name of the Mongo Database which the Spring Cloud App should be associated with. Should only be set when api_type is mongo.
	CosmosDBMongoDatabaseName *string `json:"cosmosdbMongoDatabaseName,omitempty" tf:"cosmosdb_mongo_database_name,omitempty"`

	// Specifies the name of the SQL Database which the Spring Cloud App should be associated with. Should only be set when api_type is sql.
	CosmosDBSQLDatabaseName *string `json:"cosmosdbSqlDatabaseName,omitempty" tf:"cosmosdb_sql_database_name,omitempty"`
}

type SpringCloudAppCosmosDBAssociationObservation struct {

	// Specifies the API type which should be used when connecting to the CosmosDB Account. Possible values are cassandra, gremlin, mongo, sql or table. Changing this forces a new resource to be created.
	APIType *string `json:"apiType,omitempty" tf:"api_type,omitempty"`

	// Specifies the CosmosDB Account access key.
	CosmosDBAccessKey *string `json:"cosmosdbAccessKey,omitempty" tf:"cosmosdb_access_key,omitempty"`

	// Specifies the ID of the CosmosDB Account. Changing this forces a new resource to be created.
	CosmosDBAccountID *string `json:"cosmosdbAccountId,omitempty" tf:"cosmosdb_account_id,omitempty"`

	// Specifies the name of the Cassandra Keyspace which the Spring Cloud App should be associated with. Should only be set when api_type is cassandra.
	CosmosDBCassandraKeySpaceName *string `json:"cosmosdbCassandraKeyspaceName,omitempty" tf:"cosmosdb_cassandra_keyspace_name,omitempty"`

	// Specifies the name of the Gremlin Database which the Spring Cloud App should be associated with. Should only be set when api_type is gremlin.
	CosmosDBGremlinDatabaseName *string `json:"cosmosdbGremlinDatabaseName,omitempty" tf:"cosmosdb_gremlin_database_name,omitempty"`

	// Specifies the name of the Gremlin Graph which the Spring Cloud App should be associated with. Should only be set when api_type is gremlin.
	CosmosDBGremlinGraphName *string `json:"cosmosdbGremlinGraphName,omitempty" tf:"cosmosdb_gremlin_graph_name,omitempty"`

	// Specifies the name of the Mongo Database which the Spring Cloud App should be associated with. Should only be set when api_type is mongo.
	CosmosDBMongoDatabaseName *string `json:"cosmosdbMongoDatabaseName,omitempty" tf:"cosmosdb_mongo_database_name,omitempty"`

	// Specifies the name of the SQL Database which the Spring Cloud App should be associated with. Should only be set when api_type is sql.
	CosmosDBSQLDatabaseName *string `json:"cosmosdbSqlDatabaseName,omitempty" tf:"cosmosdb_sql_database_name,omitempty"`

	// The ID of the Spring Cloud Application CosmosDB Association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the ID of the Spring Cloud Application where this Association is created. Changing this forces a new resource to be created.
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`
}

type SpringCloudAppCosmosDBAssociationParameters struct {

	// Specifies the API type which should be used when connecting to the CosmosDB Account. Possible values are cassandra, gremlin, mongo, sql or table. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	APIType *string `json:"apiType,omitempty" tf:"api_type,omitempty"`

	// Specifies the CosmosDB Account access key.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("primary_key",true)
	// +kubebuilder:validation:Optional
	CosmosDBAccessKey *string `json:"cosmosdbAccessKey,omitempty" tf:"cosmosdb_access_key,omitempty"`

	// Reference to a Account in cosmosdb to populate cosmosdbAccessKey.
	// +kubebuilder:validation:Optional
	CosmosDBAccessKeyRef *v1.Reference `json:"cosmosdbAccessKeyRef,omitempty" tf:"-"`

	// Selector for a Account in cosmosdb to populate cosmosdbAccessKey.
	// +kubebuilder:validation:Optional
	CosmosDBAccessKeySelector *v1.Selector `json:"cosmosdbAccessKeySelector,omitempty" tf:"-"`

	// Specifies the ID of the CosmosDB Account. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/cosmosdb/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	CosmosDBAccountID *string `json:"cosmosdbAccountId,omitempty" tf:"cosmosdb_account_id,omitempty"`

	// Reference to a Account in cosmosdb to populate cosmosdbAccountId.
	// +kubebuilder:validation:Optional
	CosmosDBAccountIDRef *v1.Reference `json:"cosmosdbAccountIdRef,omitempty" tf:"-"`

	// Selector for a Account in cosmosdb to populate cosmosdbAccountId.
	// +kubebuilder:validation:Optional
	CosmosDBAccountIDSelector *v1.Selector `json:"cosmosdbAccountIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Cassandra Keyspace which the Spring Cloud App should be associated with. Should only be set when api_type is cassandra.
	// +kubebuilder:validation:Optional
	CosmosDBCassandraKeySpaceName *string `json:"cosmosdbCassandraKeyspaceName,omitempty" tf:"cosmosdb_cassandra_keyspace_name,omitempty"`

	// Specifies the name of the Gremlin Database which the Spring Cloud App should be associated with. Should only be set when api_type is gremlin.
	// +kubebuilder:validation:Optional
	CosmosDBGremlinDatabaseName *string `json:"cosmosdbGremlinDatabaseName,omitempty" tf:"cosmosdb_gremlin_database_name,omitempty"`

	// Specifies the name of the Gremlin Graph which the Spring Cloud App should be associated with. Should only be set when api_type is gremlin.
	// +kubebuilder:validation:Optional
	CosmosDBGremlinGraphName *string `json:"cosmosdbGremlinGraphName,omitempty" tf:"cosmosdb_gremlin_graph_name,omitempty"`

	// Specifies the name of the Mongo Database which the Spring Cloud App should be associated with. Should only be set when api_type is mongo.
	// +kubebuilder:validation:Optional
	CosmosDBMongoDatabaseName *string `json:"cosmosdbMongoDatabaseName,omitempty" tf:"cosmosdb_mongo_database_name,omitempty"`

	// Specifies the name of the SQL Database which the Spring Cloud App should be associated with. Should only be set when api_type is sql.
	// +kubebuilder:validation:Optional
	CosmosDBSQLDatabaseName *string `json:"cosmosdbSqlDatabaseName,omitempty" tf:"cosmosdb_sql_database_name,omitempty"`

	// Specifies the ID of the Spring Cloud Application where this Association is created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/appplatform/v1beta1.SpringCloudApp
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SpringCloudAppID *string `json:"springCloudAppId,omitempty" tf:"spring_cloud_app_id,omitempty"`

	// Reference to a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDRef *v1.Reference `json:"springCloudAppIdRef,omitempty" tf:"-"`

	// Selector for a SpringCloudApp in appplatform to populate springCloudAppId.
	// +kubebuilder:validation:Optional
	SpringCloudAppIDSelector *v1.Selector `json:"springCloudAppIdSelector,omitempty" tf:"-"`
}

// SpringCloudAppCosmosDBAssociationSpec defines the desired state of SpringCloudAppCosmosDBAssociation
type SpringCloudAppCosmosDBAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SpringCloudAppCosmosDBAssociationParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SpringCloudAppCosmosDBAssociationInitParameters `json:"initProvider,omitempty"`
}

// SpringCloudAppCosmosDBAssociationStatus defines the observed state of SpringCloudAppCosmosDBAssociation.
type SpringCloudAppCosmosDBAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SpringCloudAppCosmosDBAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAppCosmosDBAssociation is the Schema for the SpringCloudAppCosmosDBAssociations API. Associates a
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SpringCloudAppCosmosDBAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.apiType) || has(self.initProvider.apiType)",message="apiType is a required parameter"
	Spec   SpringCloudAppCosmosDBAssociationSpec   `json:"spec"`
	Status SpringCloudAppCosmosDBAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SpringCloudAppCosmosDBAssociationList contains a list of SpringCloudAppCosmosDBAssociations
type SpringCloudAppCosmosDBAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpringCloudAppCosmosDBAssociation `json:"items"`
}

// Repository type metadata.
var (
	SpringCloudAppCosmosDBAssociation_Kind             = "SpringCloudAppCosmosDBAssociation"
	SpringCloudAppCosmosDBAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SpringCloudAppCosmosDBAssociation_Kind}.String()
	SpringCloudAppCosmosDBAssociation_KindAPIVersion   = SpringCloudAppCosmosDBAssociation_Kind + "." + CRDGroupVersion.String()
	SpringCloudAppCosmosDBAssociation_GroupVersionKind = CRDGroupVersion.WithKind(SpringCloudAppCosmosDBAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&SpringCloudAppCosmosDBAssociation{}, &SpringCloudAppCosmosDBAssociationList{})
}
