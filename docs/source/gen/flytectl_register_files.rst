.. _flytectl_register_files:

flytectl register files
-----------------------

Register file resources

Synopsis
~~~~~~~~



Registers all the serialized protobuf files including tasks, workflows and launchplans with default v1 version.
If previously registered entities with v1 version are present, the command will fail immediately on the first such encounter.
::

 flytectl register file  _pb_output/* -d development  -p flytesnacks
	
There is no difference between registration and fast registration. In fast registration, the input provided by the user is fast serialized proto that is generated by pyflyte. If FlyteCTL finds any source code in users' input, it considers the registration as fast registration. FlyteCTL finds input file by searching an archive file whose name starts with fast and has .tar.gz extension. When the user runs pyflyte with --fast flag then pyflyte creates serialize proto and it also creates source code archive file in the same directory. 
SourceUploadPath is an optional flag. By default, FlyteCTL will create SourceUploadPath from your storage config. In case of s3 FlyteCTL will upload code base in s3://{{DEFINE_BUCKET_IN_STORAGE_CONFIG}}/fast/{{VERSION}}-fast{{MD5_CREATED_BY_PYFLYTE}.tar.gz}. 
::

 flytectl register file  _pb_output/* -d development  -p flytesnacks  --version v2 
	
In case of fast registration, if the SourceUploadPath flag is defined, FlyteCTL will not use the default directory to upload the source code. Instead, it will override the destination path on the registration.
::

 flytectl register file  _pb_output/* -d development  -p flytesnacks  --version v2 --SourceUploadPath="s3://dummy/fast" 
	
Using archive file. Currently supported extensions are .tgz and .tar. They can be local or remote files served through http/https.
Use --archive flag:

::

  flytectl register files  http://localhost:8080/_pb_output.tar -d development  -p flytesnacks --archive

Using local tgz file:

::

 flytectl register files  _pb_output.tgz -d development  -p flytesnacks --archive

If you wish to continue executing registration on other files by ignoring the errors including the version conflicts, then send the continueOnError flag:

::

 flytectl register file  _pb_output/* -d development  -p flytesnacks --continueOnError

Using short format of continueOnError flag:
::

 flytectl register file  _pb_output/* -d development  -p flytesnacks --continueOnError

Override the default version v1 using version string:
::

 flytectl register file  _pb_output/* -d development  -p flytesnacks --version v2

Changing the o/p format has no effect on the registration. The O/p is currently available only in table format:

::

 flytectl register file  _pb_output/* -d development  -p flytesnacks --continueOnError -o yaml

Override IamRole during registration:

::

 flytectl register file  _pb_output/* -d development  -p flytesnacks --continueOnError --version v2 --assumableIamRole "arn:aws:iam::123456789:role/dummy"

Override Kubernetes service account during registration:

::

 flytectl register file  _pb_output/* -d development  -p flytesnacks --continueOnError --version v2 --k8sServiceAccount "kubernetes-service-account"

Override Output location prefix during registration:

::

 flytectl register file  _pb_output/* -d development  -p flytesnacks --continueOnError --version v2 --outputLocationPrefix "s3://dummy/prefix"

Override Destination dir of source code in container during registration:

::

 flytectl register file  _pb_output/* -d development  -p flytesnacks --continueOnError --version v2 --destinationDirectory "/root" 
	
Usage


::

  flytectl register files [flags]

Options
~~~~~~~

::

      --archive                       pass in archive file either an http link or local path.
      --assumableIamRole string        custom assumable iam auth role to register launch plans with.
      --continueOnError               continue on error when registering files.
      --destinationDirectory string    Location of source code in container.
      --dryRun                        execute command without making any modifications.
      --force                         force use of version number on entities registered with flyte.
  -h, --help                          help for files
      --k8ServiceAccount string        deprecated. Please use --K8sServiceAccount
      --k8sServiceAccount string       custom kubernetes service account auth role to register launch plans with.
      --outputLocationPrefix string    custom output location prefix for offloaded types (files/schemas).
      --sourceUploadPath string        Location for source code in storage.
      --version string                version of the entity to be registered with flyte which are un-versioned after serialization.

Options inherited from parent commands
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

::

      --admin.authorizationHeader string           Custom metadata header to pass JWT
      --admin.authorizationServerUrl string        This is the URL to your IdP's authorization server. It'll default to Endpoint
      --admin.caCertFilePath string                Use specified certificate file to verify the admin server peer.
      --admin.clientId string                      Client ID (default "flytepropeller")
      --admin.clientSecretLocation string          File containing the client secret (default "/etc/secrets/client_secret")
      --admin.command strings                      Command for external authentication token generation
      --admin.endpoint string                      For admin types,  specify where the uri of the service is located.
      --admin.insecure                             Use insecure connection.
      --admin.insecureSkipVerify                   InsecureSkipVerify controls whether a client verifies the server's certificate chain and host name. Caution : shouldn't be use for production usecases'
      --admin.maxBackoffDelay string               Max delay for grpc backoff (default "8s")
      --admin.maxRetries int                       Max number of gRPC retries (default 4)
      --admin.perRetryTimeout string               gRPC per retry timeout (default "15s")
      --admin.pkceConfig.refreshTime string         (default "5m0s")
      --admin.pkceConfig.timeout string             (default "15s")
      --admin.scopes strings                       List of scopes to request
      --admin.tokenUrl string                      OPTIONAL: Your IdP's token endpoint. It'll be discovered from flyte admin's OAuth Metadata endpoint if not provided.
      --admin.useAuth                              Deprecated: Auth will be enabled/disabled based on admin's dynamically discovered information.
  -c, --config string                              config file (default is $HOME/.flyte/config.yaml)
  -d, --domain string                              Specifies the Flyte project's domain.
      --logger.formatter.type string               Sets logging format type. (default "json")
      --logger.level int                           Sets the minimum logging level. (default 4)
      --logger.mute                                Mutes all logs regardless of severity. Intended for benchmarks/tests only.
      --logger.show-source                         Includes source code location in logs.
  -o, --output string                              Specifies the output type - supported formats [TABLE JSON YAML DOT DOTURL]. NOTE: dot, doturl are only supported for Workflow (default "TABLE")
  -p, --project string                             Specifies the Flyte project.
      --storage.cache.max_size_mbs int             Maximum size of the cache where the Blob store data is cached in-memory. If not specified or set to 0,  cache is not used
      --storage.cache.target_gc_percent int        Sets the garbage collection target percentage.
      --storage.connection.access-key string       Access key to use. Only required when authtype is set to accesskey.
      --storage.connection.auth-type string        Auth Type to use [iam, accesskey]. (default "iam")
      --storage.connection.disable-ssl             Disables SSL connection. Should only be used for development.
      --storage.connection.endpoint string         URL for storage client to connect to.
      --storage.connection.region string           Region to connect to. (default "us-east-1")
      --storage.connection.secret-key string       Secret to use when accesskey is set.
      --storage.container string                   Initial container (in s3 a bucket) to create -if it doesn't exist-.'
      --storage.defaultHttpClient.timeout string   Sets time out on the http client. (default "0s")
      --storage.enable-multicontainer              If this is true,  then the container argument is overlooked and redundant. This config will automatically open new connections to new containers/buckets as they are encountered
      --storage.limits.maxDownloadMBs int          Maximum allowed download size (in MBs) per call. (default 2)
      --storage.stow.config stringToString         Configuration for stow backend. Refer to github/graymeta/stow (default [])
      --storage.stow.kind string                   Kind of Stow backend to use. Refer to github/graymeta/stow
      --storage.type string                        Sets the type of storage to configure [s3/minio/local/mem/stow]. (default "s3")

SEE ALSO
~~~~~~~~

* :doc:`flytectl_register` 	 - Register tasks/workflows/launchplans from a list of generated serialized files.

