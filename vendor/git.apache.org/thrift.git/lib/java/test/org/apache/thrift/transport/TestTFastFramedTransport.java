/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.thrift.transport;

public class TestTFastFramedTransport extends TestTFramedTransport {
  protected final static int INITIAL_CAPACITY = 50;

  @Override
  protected TTransport getTransport(TTransport underlying) {
    return new TFastFramedTransport(underlying, INITIAL_CAPACITY, 10 * 1024 * 1024);
  }

  @Override
  protected TTransport getTransport(TTransport underlying, int maxLength) {
    return new TFastFramedTransport(underlying, INITIAL_CAPACITY, maxLength);
  }
}
