package io.fabric8.kudo.client;

import io.fabric8.kubernetes.client.Config;
import io.fabric8.kubernetes.client.dsl.base.OperationContext;

import okhttp3.OkHttpClient;

public class OperatorDependencyOperationsImplbak {

    public OperatorDependencyOperationsImplbak(OkHttpClient client, Config config) {
        this(new OperationContext().withOkhttpClient(client).withConfig(config));
    }

    public OperatorDependencyOperationsImplbak(OperationContext context) {
    }

    /*public io.fabric8.kudo.client.apis.v1beta1.internal.OperatorDependencyOperationsImplbak newInstance(OperationContext context) {
        return new io.fabric8.kudo.client.apis.v1beta1.internal.OperatorDependencyOperationsImplbak(context);
    }*/
}
